package ux

const (
	Size int  = 32 << (^uint(0) >> 63)
	Half int  = Size - 1
	Min  uint = 0
	Max  uint = 1<<Size - 1
)
