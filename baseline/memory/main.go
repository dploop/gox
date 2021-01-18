package main

import (
	"log"
	"time"
)

const (
	repeat = 100
	length = 100000000
)

func main() {
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = i
	}
	start := time.Now()
	var sum int64
	for r := 0; r < repeat; r++ {
		for i := 0; i < length; i++ {
			sum += int64(array[i])
		}
	}
	elapsed := time.Since(start)
	log.Printf("sum is %v", sum)
	ops := float64(length) / elapsed.Seconds()
	log.Printf("%.0f operation per second", ops)
}
