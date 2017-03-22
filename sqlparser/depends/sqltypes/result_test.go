// Copyright 2015, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqltypes

import (
	"fmt"
	"reflect"
	"testing"

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
)

func TestRepair(t *testing.T) {
	fields := []*querypb.Field{{
		Type: Int64,
	}, {
		Type: VarChar,
	}}
	in := Result{
		Rows: [][]Value{
			{testVal(VarBinary, "1"), testVal(VarBinary, "aa")},
			{testVal(VarBinary, "2"), testVal(VarBinary, "bb")},
		},
	}
	want := Result{
		Rows: [][]Value{
			{testVal(Int64, "1"), testVal(VarChar, "aa")},
			{testVal(Int64, "2"), testVal(VarChar, "bb")},
		},
	}
	in.Repair(fields)
	if !reflect.DeepEqual(in, want) {
		t.Errorf("Repair:\n%#v, want\n%#v", in, want)
	}
}

func TestCopy(t *testing.T) {
	in := &Result{
		Fields: []*querypb.Field{{
			Type: Int64,
		}, {
			Type: VarChar,
		}},
		InsertID:     1,
		RowsAffected: 2,
		Rows: [][]Value{
			{testVal(Int64, "1"), MakeTrusted(Null, nil)},
			{testVal(Int64, "2"), MakeTrusted(VarChar, nil)},
			{testVal(Int64, "3"), testVal(VarChar, "")},
		},
		Extras: &querypb.ResultExtras{
			EventToken: &querypb.EventToken{
				Timestamp: 123,
				Shard:     "sh",
				Position:  "po",
			},
			Fresher: true,
		},
	}
	want := &Result{
		Fields: []*querypb.Field{{
			Type: Int64,
		}, {
			Type: VarChar,
		}},
		InsertID:     1,
		RowsAffected: 2,
		Rows: [][]Value{
			{testVal(Int64, "1"), MakeTrusted(Null, nil)},
			{testVal(Int64, "2"), testVal(VarChar, "")},
			{testVal(Int64, "3"), testVal(VarChar, "")},
		},
		Extras: &querypb.ResultExtras{
			EventToken: &querypb.EventToken{
				Timestamp: 123,
				Shard:     "sh",
				Position:  "po",
			},
			Fresher: true,
		},
	}
	out := in.Copy()
	// Change in so we're sure out got actually copied
	in.Fields[0].Type = VarChar
	in.Rows[0][0] = testVal(VarChar, "aa")
	if !reflect.DeepEqual(out, want) {
		t.Errorf("Copy:\n%#v, want\n%#v", out, want)
	}
}

func TestStripFieldNames(t *testing.T) {
	testcases := []struct {
		name     string
		in       *Result
		expected *Result
	}{{
		name:     "no fields",
		in:       &Result{},
		expected: &Result{},
	}, {
		name: "empty fields",
		in: &Result{
			Fields: []*querypb.Field{},
		},
		expected: &Result{
			Fields: []*querypb.Field{},
		},
	}, {
		name: "no name",
		in: &Result{
			Fields: []*querypb.Field{{
				Type: Int64,
			}, {
				Type: VarChar,
			}},
		},
		expected: &Result{
			Fields: []*querypb.Field{{
				Type: Int64,
			}, {
				Type: VarChar,
			}},
		},
	}, {
		name: "names",
		in: &Result{
			Fields: []*querypb.Field{{
				Name: "field1",
				Type: Int64,
			}, {
				Name: "field2",
				Type: VarChar,
			}},
		},
		expected: &Result{
			Fields: []*querypb.Field{{
				Type: Int64,
			}, {
				Type: VarChar,
			}},
		},
	}}
	for _, tcase := range testcases {
		inCopy := tcase.in.Copy()
		out := inCopy.StripFieldNames()
		if !reflect.DeepEqual(out, tcase.expected) {
			t.Errorf("StripFieldNames unexpected result for %v: %v", tcase.name, out)
		}
		if len(tcase.in.Fields) > 0 {
			// check the out array is different than the in array.
			if out.Fields[0] == inCopy.Fields[0] {
				t.Errorf("StripFieldNames modified original Field for %v", tcase.name)
			}
		}
		// check we didn't change the original result.
		if !reflect.DeepEqual(tcase.in, inCopy) {
			t.Errorf("StripFieldNames modified original result")
		}
	}
}

func TestSorter(t *testing.T) {
	rs := &Result{
		Fields: []*querypb.Field{{
			Name: "user",
			Type: VarChar,
		}, {
			Name: "language",
			Type: VarChar,
		}, {
			Name: "lines",
			Type: Int64,
		},
		},
		Rows: [][]Value{
			{testVal(VarChar, "gri"), testVal(VarChar, "Go"), testVal(Int64, "100")},
			{testVal(VarChar, "ken"), testVal(VarChar, "C"), testVal(Int64, "150")},
			{testVal(VarChar, "glenda"), testVal(VarChar, "Go"), testVal(Int64, "200")},
			{testVal(VarChar, "rsc"), testVal(VarChar, "Go"), testVal(Int64, "200")},
			{testVal(VarChar, "r"), testVal(VarChar, "Go"), testVal(Int64, "200")},
			{testVal(VarChar, "ken"), testVal(VarChar, "Go"), testVal(Int64, "200")},
			{testVal(VarChar, "dmr"), testVal(VarChar, "C"), testVal(Int64, "100")},
			{testVal(VarChar, "r"), testVal(VarChar, "C"), testVal(Int64, "150")},
			{testVal(VarChar, "gri"), testVal(VarChar, "Smalltalk"), testVal(Int64, "80")},
		},
	}

	// asc
	{
		fields := []string{
			"user",
			"language",
			"lines",
		}
		wants := []string{
			"[[dmr C 100] [glenda Go 200] [gri Go 100] [gri Smalltalk 80] [ken C 150] [ken Go 200] [r Go 200] [r C 150] [rsc Go 200]]",
			"[[dmr C 100] [ken C 150] [r C 150] [glenda Go 200] [rsc Go 200] [r Go 200] [ken Go 200] [gri Go 100] [gri Smalltalk 80]]",
			"[[gri Smalltalk 80] [gri Go 100] [dmr C 100] [ken C 150] [r C 150] [rsc Go 200] [r Go 200] [ken Go 200] [glenda Go 200]]",
		}

		for i, field := range fields {
			rs1 := rs.Copy()
			rs1.OrderedByAsc(field)
			rs1.Sort()

			want := wants[i]
			got := fmt.Sprintf("%+v", rs1.Rows)
			if want != got {
				t.Errorf("want:%s\n, got:%s", want, got)
			}
		}
	}

	// desc
	{
		fields := []string{
			"user",
			"language",
			"lines",
		}
		wants := []string{
			"[[rsc Go 200] [r C 150] [r Go 200] [ken Go 200] [ken C 150] [gri Go 100] [gri Smalltalk 80] [glenda Go 200] [dmr C 100]]",
			"[[gri Smalltalk 80] [gri Go 100] [rsc Go 200] [r Go 200] [ken Go 200] [glenda Go 200] [ken C 150] [dmr C 100] [r C 150]]",
			"[[glenda Go 200] [rsc Go 200] [r Go 200] [ken Go 200] [ken C 150] [r C 150] [gri Go 100] [dmr C 100] [gri Smalltalk 80]]",
		}

		for i, field := range fields {
			rs1 := rs.Copy()
			rs1.OrderedByDesc(field)
			rs1.Sort()

			want := wants[i]
			got := fmt.Sprintf("%+v", rs1.Rows)
			if want != got {
				t.Errorf("want:%s\n, got:%s", want, got)
			}
		}
	}

	// user + language + lines asc
	{
		fields := []string{
			"user",
			"language",
			"lines",
		}

		rs1 := rs.Copy()
		for _, field := range fields {
			rs1.OrderedByAsc(field)
		}
		rs1.Sort()
		want := "[[dmr C 100] [glenda Go 200] [gri Go 100] [gri Smalltalk 80] [ken C 150] [ken Go 200] [r C 150] [r Go 200] [rsc Go 200]]"
		got := fmt.Sprintf("%+v", rs1.Rows)
		if want != got {
			t.Errorf("want:%s\n, got:%s", want, got)
		}
	}

	// user + language + lines desc
	{
		fields := []string{
			"user",
			"language",
			"lines",
		}

		rs1 := rs.Copy()
		for _, field := range fields {
			rs1.OrderedByDesc(field)
		}
		rs1.Sort()
		want := "[[rsc Go 200] [r Go 200] [r C 150] [ken Go 200] [ken C 150] [gri Smalltalk 80] [gri Go 100] [glenda Go 200] [dmr C 100]]"
		got := fmt.Sprintf("%+v", rs1.Rows)
		if want != got {
			t.Errorf("want:%s\n, got:%s", want, got)
		}
	}
}

func TestSorterType(t *testing.T) {
	rs := &Result{
		Fields: []*querypb.Field{{
			Name: "ID",
			Type: Uint24,
		}, {
			Name: "cost",
			Type: Float32,
		},
		},
		Rows: [][]Value{
			{testVal(Uint24, "3"), testVal(Float32, "3.1415926")},
			{testVal(Uint24, "7"), testVal(Float32, "3.1415926")},
			{testVal(Uint24, "2"), testVal(Float32, "3.1415927")},
		},
	}

	// asc
	{
		fields := []string{
			"ID",
			"cost",
		}
		wants := []string{
			"[[2 3.1415927] [3 3.1415926] [7 3.1415926]]",
			"[[3 3.1415926] [7 3.1415926] [2 3.1415927]]",
		}

		for i, field := range fields {
			rs1 := rs.Copy()
			rs1.OrderedByAsc(field)
			rs1.Sort()

			want := wants[i]
			got := fmt.Sprintf("%+v", rs1.Rows)
			if want != got {
				t.Errorf("want:%s\n, got:%s", want, got)
			}
		}
	}

	// desc
	{
		fields := []string{
			"ID",
			"cost",
		}
		wants := []string{
			"[[7 3.1415926] [3 3.1415926] [2 3.1415927]]",
			"[[2 3.1415927] [3 3.1415926] [7 3.1415926]]",
		}

		for i, field := range fields {
			rs1 := rs.Copy()
			rs1.OrderedByDesc(field)
			rs1.Sort()

			want := wants[i]
			got := fmt.Sprintf("%+v", rs1.Rows)
			if want != got {
				t.Errorf("want:%s\n, got:%s", want, got)
			}
		}
	}
}

func TestSorterError(t *testing.T) {
	rs := &Result{
		Fields: []*querypb.Field{{
			Name: "ID",
			Type: Uint24,
		}, {
			Name: "cost",
			Type: Float32,
		},
		},
		Rows: [][]Value{
			{testVal(Uint24, "3"), testVal(Float32, "3.1415926")},
			{testVal(Uint24, "7"), testVal(Float32, "3.1415926")},
			{testVal(Uint24, "2"), testVal(Float32, "3.1415927")},
		},
	}

	// Field error.
	{
		{
			rs1 := rs.Copy()
			err := rs1.OrderedByAsc("xx")
			want := "can.not.find.the.orderby.field[xx].direction.asc"
			got := err.Error()
			if want != got {
				t.Errorf("want:%s\n, got:%s", want, got)
			}
		}

		{
			rs1 := rs.Copy()
			err := rs1.OrderedByDesc("xx")
			want := "can.not.find.the.orderby.field[xx].direction.desc"
			got := err.Error()
			if want != got {
				t.Errorf("want:%s\n, got:%s", want, got)
			}
		}

	}
}
