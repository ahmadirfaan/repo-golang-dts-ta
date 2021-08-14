package main

import "fmt"

func main() {
	// Menggunakan Make
	a := make([]int, 3)
	a = []int{1,2,3}
	// Buat slice of int bernama "a" dengan panjang 3, {1, 2,3}
	printSlice("a", a)

	var b = make([]int, 5)
	b = []int{1,2,3,4,5}
	// Buat slice of int bernama "b" dengan panjang 5 {1, 2, 3, 4, 5}
	printSlice("b", b)

	var c = b[:2]
	// Buat variabel c dengan dua data awal b
	printSlice("c", c)
	
	var d = b[1:5]
	// Buat variabel d dengan data ke 2 sampai 4
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s, len(x), cap(x), x)
}