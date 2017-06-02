// Copyright 2015, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Copyright (c) XeLabs
// BohuTANG

package sqltypes

import (
	"bytes"
)

func Operator(v1 Value, v2 Value, fn func(x interface{}, y interface{}) interface{}) Value {
	// Sum field type is Decimal, we convert it to golang Float64.
	switch v1.Type() {
	case Decimal:
		v1 = MakeTrusted(Float64, v1.Raw())
	}
	switch v2.Type() {
	case Decimal:
		v2 = MakeTrusted(Float64, v2.Raw())
	}

	v1n := v1.ToNative()
	v2n := v2.ToNative()
	val := fn(v1n, v2n)
	v, err := BuildValue(val)
	if err != nil {
		panic(err)
	}
	return v
}

func SumFn(x interface{}, y interface{}) interface{} {
	var v interface{}
	switch x.(type) {
	case int64:
		v = (x.(int64) + y.(int64))
	case uint64:
		v = (x.(uint64) + y.(uint64))
	case float64:
		v = (x.(float64) + y.(float64))
	case []uint8: // We only support numerical value sum.
		v = 0
	}
	return v
}

func MinFn(x interface{}, y interface{}) interface{} {
	v := x
	switch x.(type) {
	case int64:
		if x.(int64) > y.(int64) {
			v = y
		}
	case uint64:
		if x.(uint64) > y.(uint64) {
			v = y
		}
	case float64:
		if x.(float64) > y.(float64) {
			v = y
		}
	case []uint8:
		if bytes.Compare(x.([]uint8), y.([]uint8)) > 0 {
			v = y
		}
	}
	return v
}

func MaxFn(x interface{}, y interface{}) interface{} {
	v := x
	switch x.(type) {
	case int64:
		if x.(int64) < y.(int64) {
			v = y
		}
	case uint64:
		if x.(uint64) < y.(uint64) {
			v = y
		}
	case float64:
		if x.(float64) < y.(float64) {
			v = y
		}
	case []uint8:
		if bytes.Compare(x.([]uint8), y.([]uint8)) < 0 {
			v = y
		}
	}
	return v
}

func DivFn(x interface{}, y interface{}) interface{} {
	var v1, v2 float64

	switch x.(type) {
	case int64:
		v1 = float64(x.(int64))
	case uint64:
		v1 = float64(x.(uint64))
	case float64:
		v1 = x.(float64)
	case []uint8: // We only support numerical value div.
		return 0
	}
	switch y.(type) {
	case int64:
		v2 = float64(y.(int64))
	case uint64:
		v2 = float64(y.(uint64))
	case float64:
		v2 = y.(float64)
	case []uint8: // We only support numerical value div.
		return 0
	}
	return v1 / v2
}
