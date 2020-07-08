package main

import (
	"log"
	"strings"

	"github.com/dploop/gox/utils"
)

// https://go101.org/article/bounds-check-elimination.html
func main() {
	x := strings.Repeat("hello", 999) + "world!"
	y := strings.Repeat("hello", 998) + "world!"
	slower(x, y)
	faster(x, y)
}

func slower(x string, y string) {
	if len(x) > len(y) {
		x, y = y, x
	}
	defer utils.LogElapsed("slower")()
	var sum int
	for k := 0; k < 1000000; k++ {
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				break
			}
			sum++
		}
	}
	log.Printf("slower sum is %v", sum)
}

func faster(x string, y string) {
	if len(x) > len(y) {
		x, y = y, x
	}
	if len(x) > len(y) {
		return
	}
	defer utils.LogElapsed("faster")()
	var sum int
	for k := 0; k < 1000000; k++ {
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				break
			}
			sum++
		}
	}
	log.Printf("faster sum is %v", sum)
}
