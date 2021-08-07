package main

import (
	"fmt"
	"math/rand"
)

// merupakan package tersendiri di golang
func main() {
	if n := rand.Intn(100); n == 0 {
		fmt.Println("That's to low")
	} else if n > 5 {
		fmt.Println("That's too Big", n)
	} else {
		fmt.Println("That's a good number: ", n)
	}
}
