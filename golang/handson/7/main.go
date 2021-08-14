package main

import "fmt"

func fibonacci(until int) {
	n1 := 0
	n2 := 1
	fmt.Print(n1, " ", n2)
	for i := 2; i < until; i++ {
		n3 := n1 + n2
		fmt.Print(" ", n3)
		n1 = n2
		n2 = n3
	}
	
	// buatlah fungsingya
}

func main() {
	fibonacci(20)
}
