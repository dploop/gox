package main

import (
	"log"
	"sync"
	"time"
)

const (
	concurrency = 2
	repeat      = 100000000
	modulo      = 16
)

func main() {
	var (
		sum int64
		mu  sync.Mutex
	)
	start := time.Now()
	var wg sync.WaitGroup
	for c := 0; c < concurrency; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := 0; r < repeat; r++ {
				d := int64(r % modulo)
				mu.Lock()
				sum += d
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("sum is %v", sum)
	ops := float64(repeat*concurrency) / elapsed.Seconds()
	log.Printf("%.0f operation per second", ops)
}
