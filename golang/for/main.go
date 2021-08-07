package main

import "fmt"

func main() {
	// C Style Loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Style Golang
	i := 1
	for i < 100 {
		fmt.Println(i)
		i *= 2
	}

	//For Range Statement
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

}
