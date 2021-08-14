package main

import "fmt"

func main() {
	// buat variabel m dengan map key: string dan value: int
	m := map[string]int{}
	m["Answer"] = 42
	// beri nilai "Answer" dengan 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	// ganti nilai "Answer" dengan 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	// hapus "Answer
	fmt.Println("The value:", m["Answer"])
	var v, ok = m["Answer"]
	// gunakan pengecekan: elem, ok = m[key]
	fmt.Println("The value:", v, "Present?", ok)
}