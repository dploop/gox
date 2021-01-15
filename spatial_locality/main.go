package main

import (
	"log"

	"github.com/dploop/gox/utils"
)

const (
	repeat   = 100
	length   = 9999999
	slowStep = 1000
	fastStep = 9999998
)

func main() {
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = i
	}
	do("slow", array, slowStep)
	do("fast", array, fastStep)
}

func do(name string, array []int, step int) {
	defer utils.LogElapsed(name)()
	var sum int64
	for r := 0; r < repeat; r++ {
		for i := step; i != 0; {
			sum += int64(array[i])
			i = (i + step) % length
		}
	}
	log.Printf("%s sum is %v", name, sum)
}
