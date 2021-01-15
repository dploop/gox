package main

import (
	"log"
	"math/rand"
	"sort"

	"github.com/dploop/gox/utils"
)

const (
	repeat = 1000
	length = 1000000
	limit  = 2000
	half   = 1000
)

func main() {
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = rand.Intn(limit) - half
	}
	do("slow", array)
	sort.Ints(array)
	do("fast", array)
}

func do(name string, array []int) {
	defer utils.LogElapsed(name)()
	var sum int
	for r := 0; r < repeat; r++ {
		for i := 0; i < length; i++ {
			if array[i] > 0 {
				sum += array[i]
			}
		}
	}
	log.Printf("%s sum is %v", name, sum)
}
