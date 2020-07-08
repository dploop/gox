package gox

import (
	"math/bits"
)

func MaxInt32Fast(x int32, y int32) int32 {
	bits.OnesCount32()
	return y & ((x - y) >> 31) | x & (^(x - y) >> 31);
}

func MaxInt32Base(x int32, y int32) int32 {
	if x < y {
		return y
	}
	return x
}
