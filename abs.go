package gox

import (
	"github.com/dploop/gox/integer/i0"
	"github.com/dploop/gox/integer/i16"
	"github.com/dploop/gox/integer/i32"
	"github.com/dploop/gox/integer/i64"
	"github.com/dploop/gox/integer/i8"
)

func AbsInt8Fast(x int8) int8 {
	s := x >> i8.Half
	return (x ^ s) - s
}

func AbsInt8Base(x int8) int8 {
	if x < 0 {
		return -x
	}
	return x
}

func AbsInt16Fast(x int16) int16 {
	s := x >> i16.Half
	return (x ^ s) - s
}

func AbsInt16Base(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}

func AbsInt32Fast(x int32) int32 {
	s := x >> i32.Half
	return (x ^ s) - s
}

func AbsInt32Base(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func AbsInt64Fast(x int64) int64 {
	s := x >> i64.Half
	return (x ^ s) - s
}

func AbsInt64Base(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func AbsIntFast(x int) int {
	s := x >> i0.Half
	return (x ^ s) - s
}

func AbsIntBase(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
