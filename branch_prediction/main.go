package main

import (
	"log"
	"math/rand"
	"sort"

	"github.com/dploop/gox/utils"
)

// https://stackoverflow.com/questions/11227809/why-is-processing-a-sorted-array-faster-than-processing-an-unsorted-array
func main() {
	array := make([]int, 1000000)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(65536) - 32768
	}
	slower(array)
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	faster(array)
}

func slower(array []int) {
	defer utils.LogElapsed("slower")()
	sum := do(array)
	log.Printf("slower sum is %v", sum)
}

func faster(array []int) {
	defer utils.LogElapsed("faster")()
	sum := do(array)
	log.Printf("faster sum is %v", sum)
}

func do(array []int) int {
	var sum int
	for k := 0; k < 1000; k++ {
		for i := 0; i < len(array); i++ {
			if array[i] > 0 {
				sum += array[i]
			}
		}
	}
	return sum
}
