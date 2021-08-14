package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// panggil "say("world)" dengan goroutine baru
	go say("world")
	// panggil "say("hello")" seperti biasa
	say("hello")
}	