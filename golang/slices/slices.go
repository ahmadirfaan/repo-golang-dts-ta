package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon", "tomato"}
	var vegetables = []string{"cucumber", "spinach", "kale", "lettuce", "tomato?"}

	fmt.Println(fruits[0])
	fmt.Println(vegetables[1])

	likedFruit := fruits[0:3]
	likedVegetables := vegetables[2:4]
	fmt.Println(likedFruit)
	fmt.Println(likedVegetables)
}
