package ix

const (
	Size int = 32 << (^uint(0) >> 63)
	Half int = Size - 1
	Min  int = -1 << Half
	Max  int = 1<<Half - 1
)
