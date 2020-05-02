package i0

const (
	Size = 32 << (^uint(0) >> 63)
	Half = Size - 1
	Min  = -1 << Half
	Max  = 1<<Half - 1
)
