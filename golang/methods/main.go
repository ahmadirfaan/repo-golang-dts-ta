package main

import "fmt"

type Square struct {S float64}

//traditonal function
func Area(s Square) float64 {
	return s.S * s.S
}

//same thing but as a method of the Point Type
func (s Square) Area() float64{
	return s.S * s.S
}

func (s *Square) ScaleBy(factor float64) {
    s.S *= factor
}

func main() {
	s1 := Square{3}
	s2 := Square{4}
	fmt.Println(Area(s1))
	fmt.Println(s2.Area())
	
	s3 := Square{2}
	s3.ScaleBy(3)
	fmt.Println(s3)
}