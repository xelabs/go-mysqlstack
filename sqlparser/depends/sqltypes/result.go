// Copyright 2015, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqltypes

import (
	"bytes"
	"fmt"
	"sort"

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
)

// Result represents a query result.
type Result struct {
	Fields       []*querypb.Field      `json:"fields"`
	RowsAffected uint64                `json:"rows_affected"`
	InsertID     uint64                `json:"insert_id"`
	Rows         [][]Value             `json:"rows"`
	Extras       *querypb.ResultExtras `json:"extras"`
	sorters      []*sorter
}

// ResultStream is an interface for receiving Result. It is used for
// RPC interfaces.
type ResultStream interface {
	// Recv returns the next result on the stream.
	// It will return io.EOF if the stream ended.
	Recv() (*Result, error)
}

// Repair fixes the type info in the rows
// to conform to the supplied field types.
func (result *Result) Repair(fields []*querypb.Field) {
	// Usage of j is intentional.
	for j, f := range fields {
		for _, r := range result.Rows {
			if r[j].typ != Null {
				r[j].typ = f.Type
			}
		}
	}
}

// Copy creates a deep copy of Result.
func (result *Result) Copy() *Result {
	out := &Result{
		InsertID:     result.InsertID,
		RowsAffected: result.RowsAffected,
	}
	if result.Fields != nil {
		fieldsp := make([]*querypb.Field, len(result.Fields))
		fields := make([]querypb.Field, len(result.Fields))
		for i, f := range result.Fields {
			fields[i] = *f
			fieldsp[i] = &fields[i]
		}
		out.Fields = fieldsp
	}
	if result.Rows != nil {
		rows := make([][]Value, len(result.Rows))
		for i, r := range result.Rows {
			rows[i] = make([]Value, len(r))
			totalLen := 0
			for _, c := range r {
				totalLen += len(c.val)
			}
			arena := make([]byte, 0, totalLen)
			for j, c := range r {
				start := len(arena)
				arena = append(arena, c.val...)
				rows[i][j] = MakeTrusted(c.typ, arena[start:start+len(c.val)])
			}
		}
		out.Rows = rows
	}
	if result.Extras != nil {
		out.Extras = &querypb.ResultExtras{
			Fresher: result.Extras.Fresher,
		}
		if result.Extras.EventToken != nil {
			out.Extras.EventToken = &querypb.EventToken{
				Timestamp: result.Extras.EventToken.Timestamp,
				Shard:     result.Extras.EventToken.Shard,
				Position:  result.Extras.EventToken.Position,
			}
		}
	}
	return out
}

// MakeRowTrusted converts a *querypb.Row to []Value based on the types
// in fields. It does not sanity check the values against the type.
// Every place this function is called, a comment is needed that explains
// why it's justified.
func MakeRowTrusted(fields []*querypb.Field, row *querypb.Row) []Value {
	sqlRow := make([]Value, len(row.Lengths))
	var offset int64
	for i, length := range row.Lengths {
		if length < 0 {
			continue
		}
		sqlRow[i] = MakeTrusted(fields[i].Type, row.Values[offset:offset+length])
		offset += length
	}
	return sqlRow
}

// StripFieldNames will return a new Result that has the same Rows,
// but the Field objects will have their Name emptied.  Note we don't
// proto.Copy each Field for performance reasons, but we only copy the
// individual fields.
func (result *Result) StripFieldNames() *Result {
	if len(result.Fields) == 0 {
		return result
	}
	r := *result
	r.Fields = make([]*querypb.Field, len(result.Fields))
	newFieldsArray := make([]querypb.Field, len(result.Fields))
	for i, f := range result.Fields {
		r.Fields[i] = &newFieldsArray[i]
		newFieldsArray[i].Type = f.Type
	}
	return &r
}

// AppendResult will combine the Results Objects of one result
// to another result.Note currently it doesn't handle cases like
// if two results have different fields.We will enhance this function.
func (result *Result) AppendResult(src *Result) {
	if src.RowsAffected == 0 && len(src.Fields) == 0 {
		return
	}
	if result.Fields == nil {
		result.Fields = src.Fields
	}
	result.RowsAffected += src.RowsAffected
	if src.InsertID != 0 {
		result.InsertID = src.InsertID
	}
	if len(result.Rows) == 0 {
		// we haven't gotten any result yet, just save the new extras.
		result.Extras = src.Extras
	} else {
		// Merge the EventTokens / Fresher flags within Extras.
		if src.Extras == nil {
			// We didn't get any from innerq. Have to clear any
			// we'd have gotten already.
			if result.Extras != nil {
				result.Extras.EventToken = nil
				result.Extras.Fresher = false
			}
		} else {
			// We may have gotten an EventToken from
			// innerqr.  If we also got one earlier, merge
			// it. If we didn't get one earlier, we
			// discard the new one.
			if result.Extras != nil {
				// Note if any of the two is nil, we get nil.
				//result.Extras.EventToken = eventtoken.Minimum(result.Extras.EventToken, src.Extras.EventToken)

				//result.Extras.Fresher = result.Extras.Fresher && src.Extras.Fresher
			}
		}
	}
	result.Rows = append(result.Rows, src.Rows...)
}

// Len is part of sort.Interface.
func (result *Result) Len() int {
	return len(result.Rows)
}

// Swap is part of sort.Interface.
func (result *Result) Swap(i, j int) {
	result.Rows[i], result.Rows[j] = result.Rows[j], result.Rows[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that is either Less or
// !Less. Note that it can call the less functions twice per call. We
// could change the functions to return -1, 0, 1 and reduce the
// number of calls for greater efficiency: an exercise for the reader.
func (result *Result) Less(i, j int) bool {
	p, q := result.Rows[i], result.Rows[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(result.sorters)-1; k++ {
		ser := result.sorters[k]
		switch {
		case ser.less(ser.idx, p, q):
			// p < q, so we have a decision.
			return true
		case ser.less(ser.idx, q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	ser := result.sorters[k]
	return ser.less(ser.idx, p, q)
}

// sorter
type LessFunc func(idx int, v1, v2 []Value) bool
type sorter struct {
	idx  int
	less LessFunc
}

func lessAscFn(idx int, v1, v2 []Value) bool {
	vn1 := v1[idx].ToNative()
	vn2 := v2[idx].ToNative()
	switch vn1.(type) {
	case int64:
		return vn1.(int64) < vn2.(int64)
	case uint64:
		return vn1.(uint64) < vn2.(uint64)
	case float64:
		return vn1.(float64) < vn2.(float64)
	case []byte:
		return bytes.Compare(vn1.([]byte), vn2.([]byte)) < 0
	default:
		panic(fmt.Sprintf("unsupported.orderby.type:%T", vn1))
	}
}

// OrderedByAsc adds a 'order by asc' operator to the result.
func (result *Result) OrderedByAsc(field string) error {
	idx := -1
	for k, f := range result.Fields {
		if f.Name == field {
			idx = k
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("can.not.find.the.orderby.field[%s].direction.asc", field)
	}
	ser := &sorter{idx: idx, less: lessAscFn}
	result.sorters = append(result.sorters, ser)
	return nil
}

func lessDescFn(idx int, v1, v2 []Value) bool {
	vn1 := v1[idx].ToNative()
	vn2 := v2[idx].ToNative()
	switch vn1.(type) {
	case int64:
		return vn2.(int64) < vn1.(int64)
	case uint64:
		return vn2.(uint64) < vn1.(uint64)
	case float64:
		return vn2.(float64) < vn1.(float64)
	case []byte:
		return bytes.Compare(vn2.([]byte), vn1.([]byte)) < 0
	default:
		panic(fmt.Sprintf("unsupport.orderby.type:%T", vn1))
	}
}

// OrderedByDesc adds a 'order by desc' operator to the result.
func (result *Result) OrderedByDesc(field string) error {
	idx := -1
	for k, f := range result.Fields {
		if f.Name == field {
			idx = k
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("can.not.find.the.orderby.field[%s].direction.desc", field)
	}
	ser := &sorter{idx: idx, less: lessDescFn}
	result.sorters = append(result.sorters, ser)
	return nil
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (result *Result) Sort() {
	sort.Sort(result)
}
