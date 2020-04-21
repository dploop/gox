package u64

func Max(x uint64, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}

func Min(x uint64, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}
