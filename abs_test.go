package gox_test

import (
	"runtime"
	"testing"

	"github.com/dploop/gox"
	"github.com/dploop/gox/integer/i16"
	"github.com/dploop/gox/integer/i32"
	"github.com/dploop/gox/integer/i64"
	"github.com/dploop/gox/integer/i8"
	"github.com/dploop/gox/integer/ix"
	"github.com/stretchr/testify/assert"
)

func TestAbsInt8(t *testing.T) {
	itemList := []struct {
		x        int8
		expected int8
	}{
		{i8.Min, i8.Min},
		{i8.Min + 1, i8.Max},
		{-42, 42},
		{-1, 1},
		{0, 0},
		{1, 1},
		{42, 42},
		{i8.Max, i8.Max},
	}
	for _, item := range itemList {
		actual := gox.AbsInt8Fast(item.x)
		assert.Equal(t, item.expected, actual)
	}
	for _, item := range itemList {
		actual := gox.AbsInt8Base(item.x)
		assert.Equal(t, item.expected, actual)
	}
}

func BenchmarkAbsInt8Fast(b *testing.B) {
	var total int8
	for i := 0; i < b.N; i++ {
		x := int8(i%20) - 10
		total ^= gox.AbsInt8Fast(x)
	}
	runtime.KeepAlive(total)
}

func BenchmarkAbsInt8Base(b *testing.B) {
	var total int8
	for i := 0; i < b.N; i++ {
		x := int8(i%20) - 10
		total ^= gox.AbsInt8Base(x)
	}
	runtime.KeepAlive(total)
}

func TestAbsInt16(t *testing.T) {
	itemList := []struct {
		x        int16
		expected int16
	}{
		{i16.Min, i16.Min},
		{i16.Min + 1, i16.Max},
		{-42, 42},
		{-1, 1},
		{0, 0},
		{1, 1},
		{42, 42},
		{i16.Max, i16.Max},
	}
	for _, item := range itemList {
		actual := gox.AbsInt16Fast(item.x)
		assert.Equal(t, item.expected, actual)
	}
	for _, item := range itemList {
		actual := gox.AbsInt16Base(item.x)
		assert.Equal(t, item.expected, actual)
	}
}

func BenchmarkAbsInt16Fast(b *testing.B) {
	var total int16
	for i := 0; i < b.N; i++ {
		x := int16(i%20) - 10
		total ^= gox.AbsInt16Fast(x)
	}
	runtime.KeepAlive(total)
}

func BenchmarkAbsInt16Base(b *testing.B) {
	var total int16
	for i := 0; i < b.N; i++ {
		x := int16(i%20) - 10
		total ^= gox.AbsInt16Base(x)
	}
	runtime.KeepAlive(total)
}

func TestAbsInt32(t *testing.T) {
	itemList := []struct {
		x        int32
		expected int32
	}{
		{i32.Min, i32.Min},
		{i32.Min + 1, i32.Max},
		{-42, 42},
		{-1, 1},
		{0, 0},
		{1, 1},
		{42, 42},
		{i32.Max, i32.Max},
	}
	for _, item := range itemList {
		actual := gox.AbsInt32Fast(item.x)
		assert.Equal(t, item.expected, actual)
	}
	for _, item := range itemList {
		actual := gox.AbsInt32Base(item.x)
		assert.Equal(t, item.expected, actual)
	}
}

func BenchmarkAbsInt32Fast(b *testing.B) {
	var total int32
	for i := 0; i < b.N; i++ {
		x := int32(i%20) - 10
		total ^= gox.AbsInt32Fast(x)
	}
	runtime.KeepAlive(total)
}

func BenchmarkAbsInt32Base(b *testing.B) {
	var total int32
	for i := 0; i < b.N; i++ {
		x := int32(i%20) - 10
		total ^= gox.AbsInt32Base(x)
	}
	runtime.KeepAlive(total)
}

func TestAbsInt64(t *testing.T) {
	itemList := []struct {
		x        int64
		expected int64
	}{
		{i64.Min, i64.Min},
		{i64.Min + 1, i64.Max},
		{-42, 42},
		{-1, 1},
		{0, 0},
		{1, 1},
		{42, 42},
		{i64.Max, i64.Max},
	}
	for _, item := range itemList {
		actual := gox.AbsInt64Fast(item.x)
		assert.Equal(t, item.expected, actual)
	}
	for _, item := range itemList {
		actual := gox.AbsInt64Base(item.x)
		assert.Equal(t, item.expected, actual)
	}
}

func BenchmarkAbsInt64Fast(b *testing.B) {
	var total int64
	for i := 0; i < b.N; i++ {
		x := int64(i%20) - 10
		total ^= gox.AbsInt64Fast(x)
	}
	runtime.KeepAlive(total)
}

func BenchmarkAbsInt64Base(b *testing.B) {
	var total int64
	for i := 0; i < b.N; i++ {
		x := int64(i%20) - 10
		total ^= gox.AbsInt64Base(x)
	}
	runtime.KeepAlive(total)
}

func TestAbsInt(t *testing.T) {
	itemList := []struct {
		x        int
		expected int
	}{
		{ix.Min, ix.Min},
		{ix.Min + 1, ix.Max},
		{-42, 42},
		{-1, 1},
		{0, 0},
		{1, 1},
		{42, 42},
		{ix.Max, ix.Max},
	}
	for _, item := range itemList {
		actual := gox.AbsIntFast(item.x)
		assert.Equal(t, item.expected, actual)
	}
	for _, item := range itemList {
		actual := gox.AbsIntBase(item.x)
		assert.Equal(t, item.expected, actual)
	}
}

func BenchmarkAbsIntFast(b *testing.B) {
	var total int
	for i := 0; i < b.N; i++ {
		x := i%20 - 10
		total ^= gox.AbsIntFast(x)
	}
	runtime.KeepAlive(total)
}

func BenchmarkAbsIntBase(b *testing.B) {
	var total int
	for i := 0; i < b.N; i++ {
		x := i%20 - 10
		total ^= gox.AbsIntBase(x)
	}
	runtime.KeepAlive(total)
}
