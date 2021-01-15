package main

import (
	"log"
	"sync"
	"sync/atomic"

	"github.com/dploop/gox/utils"
)

const (
	repeat          = 100
	length          = 100000000
	slowConcurrency = 1
	fastConcurrency = 2
)

func main() {
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = i
	}
	do("slow", array, slowConcurrency)
	do("fast", array, fastConcurrency)
}

func do(name string, array []int, concurrency int) {
	defer utils.LogElapsed(name)()
	var sum int64
	for r := 0; r < repeat; r += concurrency {
		var wg sync.WaitGroup
		for c := 0; c < concurrency; c++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var acc int64
				for i := 0; i < length; i++ {
					acc += int64(array[i])
				}
				_ = atomic.AddInt64(&sum, acc)
			}()
		}
		wg.Wait()
	}
	log.Printf("%s sum is %v", name, atomic.LoadInt64(&sum))
}
