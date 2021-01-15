package main

import (
	"log"
	"time"
)

const (
	repeat = 10000000000
	modulo = 16
)

func main() {
	var sum int64
	start := time.Now()
	for r := 0; r < repeat; r++ {
		d := int64(r % modulo)
		sum += d
	}
	elapsed := time.Since(start)
	log.Printf("sum is %v", sum)
	ops := float64(repeat) / elapsed.Seconds()
	log.Printf("%.0f operation per second", ops)
}
