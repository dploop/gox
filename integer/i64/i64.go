package i64

func Max(x int64, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func Min(x int64, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
