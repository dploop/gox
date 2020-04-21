package u32

func Max(x uint32, y uint32) uint32 {
	if x > y {
		return x
	}
	return y
}

func Min(x uint32, y uint32) uint32 {
	if x < y {
		return x
	}
	return y
}
