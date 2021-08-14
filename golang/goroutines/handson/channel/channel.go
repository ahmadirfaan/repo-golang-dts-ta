package main

import (
	"fmt"
)
  
func myfunc(ch chan int) {
    fmt.Println(234 + <-ch)
}
func main() {
    fmt.Println("start Main method")
    // Buatlah Channel
	cobaChannel := make(chan int)
    // Panggila myfunc dengan goroutine
	go myfunc(cobaChannel)
    // Masukkan angka 23 ke dalam channel
	cobaChannel <- 23

    fmt.Println("End Main method")
}