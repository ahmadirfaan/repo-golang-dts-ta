package main

import (
	"fmt"
	"strings"
)

func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}

func main() {
	var ans = Split("Irfaan-Ganteng", "-")
	v := ans[0]
	fmt.Println(ans)
	if v != ans[0] {
		fmt.Println("Error gaes")
	}
	// for i, v := range ans {
	// 	fmt.Println(v)

	// }
}
