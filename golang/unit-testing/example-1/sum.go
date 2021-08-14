package main

func Sum(a,b int) int {
	return a + b
}

func SumVariadic(numbers ...int) (result int) {
	for _, n := range numbers {
		result += n
	}
	return result
}