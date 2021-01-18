package main

import (
	"log"
	"time"
)

const (
	repeat = 10000000000
)

func main() {
	start := time.Now()
	var sum int64
	for r := 0; r < repeat; r++ {
		sum += int64(r & 1)
	}
	elapsed := time.Since(start)
	log.Printf("sum is %v", sum)
	ops := float64(repeat) / elapsed.Seconds()
	log.Printf("%.0f operation per second", ops)
}
