package main

import "fmt"

func main() {
	var arrayHandsOn [2]string
	arrayHandsOn[0] = "Go"
	arrayHandsOn[1] = "C"
	fmt.Println(arrayHandsOn[0])
	fmt.Println(arrayHandsOn[1])

	//Buat array primes
	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}