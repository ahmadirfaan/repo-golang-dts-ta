package main

import (
	"fmt"
	"math"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// buatlah for dengan range untuk pow
	// yang menghasilkan:
	for i := 0; i < len(pow); i++ {
		fmt.Println(math.Pow(2,float64(i)))
	}
	/*
		2**0 = 1
		2**1 = 2
		2**2 = 4
		2**3 = 8
		2**4 = 16
		2**5 = 32
		2**6 = 64
		2**7 = 128
	*/

	// gunakan bantuan fmt.Printf
}
