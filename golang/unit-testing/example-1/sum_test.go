package main

import (
	"fmt"
	"testing"
)

//Buatlah Test Sum
func TestSum(t *testing.T) {
	ans := Sum(2, -2)
	if ans != 0 {
		t.Errorf("Sum(2,-2) = %d; want 0", ans)
	}
}

// Buat Test Sum Variadic
func TestIntSumVariadic(t *testing.T) {
	ans := SumVariadic(2, 2, 3, 5, 6, 7)
	if ans != 25 {
		t.Errorf("IntMin(2,-2) = %d; want -2", ans)
	}
}

//Buatlah Test Sum Table Driven
func TestIntSumTableDriven(t *testing.T) {
	var tables = []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, tt := range tables {

		testname := fmt.Sprintf("%d,%d", tt.x, tt.y)
		t.Run(testname, func(t *testing.T) {
			ans := Sum(tt.x, tt.y)
			if ans != tt.n {
				t.Errorf("got %d, want %d", ans, tt.n)
			}
		})
	}
}
