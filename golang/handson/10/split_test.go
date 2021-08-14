package main

import (
	"reflect"
	"testing"
)

//Buatlah Test Split
func TestSplit(t *testing.T) {
	ans := Split("Test-Split", "-")
	var result [2]string
	result[0] = "Test"
	result[1] = "Split"
	if reflect.DeepEqual(ans, result) {
		t.Errorf("Salah Error")
	}

}
