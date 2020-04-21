package i32

func Max(x int32, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func Min(x int32, y int32) int32 {
	if x < y {
		return x
	}
	return y
}
