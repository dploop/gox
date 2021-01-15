package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

const (
	concurrency = 2
	repeat      = 100000000
	modulo      = 16
)

func main() {
	var sum int64
	start := time.Now()
	var wg sync.WaitGroup
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := 0; r < repeat; r++ {
				d := int64(r % modulo)
				atomic.AddInt64(&sum, d)
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("sum is %v", atomic.LoadInt64(&sum))
	ops := float64(repeat*concurrency) / elapsed.Seconds()
	log.Printf("%.0f operation per second", ops)
}
