package main

import (
	"log"
	"strings"

	"github.com/dploop/gox/utils"
)

const (
	repeat = 1000000
	xcount = 999
	ycount = 998
)

func main() {
	x := strings.Repeat("hello", xcount) + "world!"
	y := strings.Repeat("hello", ycount) + "world!"
	slow(x, y)
	fast(x, y)
}

func slow(x string, y string) {
	if len(x) > len(y) {
		x, y = y, x
	}
	defer utils.LogElapsed("slow")()
	var sum int
	for r := 0; r < repeat; r++ {
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				break
			}
			sum++
		}
	}
	log.Printf("slow sum is %v", sum)
}

func fast(x string, y string) {
	if len(x) > len(y) {
		x, y = y, x
	}
	if len(x) > len(y) {
		return
	}
	defer utils.LogElapsed("fast")()
	var sum int
	for k := 0; k < repeat; k++ {
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				break
			}
			sum++
		}
	}
	log.Printf("fast sum is %v", sum)
}
