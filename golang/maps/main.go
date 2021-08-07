package main

import "fmt"

func main() {
	// ages := make(map[string]int) // mapping from strings to ints

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	fmt.Println(ages)

	// Equal to

	// ages := make(map[string]int) // or map[string]int{}
	// ages["alice"] = 31
	// ages["charlie"] = 34
	// ages["alice"] = 32
	// fmt.Println(ages["alice"]) // "32"

	// delete(ages, "alice") // remove element ages["alice"]

}
