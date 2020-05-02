package u0

const (
	Size = 32 << (^uint(0) >> 63)
	Half = Size - 1
	Min  = 0
	Max  = 1<<Size - 1
)
