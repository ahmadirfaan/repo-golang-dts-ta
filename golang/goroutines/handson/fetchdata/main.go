package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Step 1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 2:", url)
	defer response.Body.Close()

	fmt.Println("Step 3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 4: ", len(body))

}

func main() {
	ch := make(chan string, 3)
	go responseSize("https://golangprograms.com", ch)
	go responseSize("https://coderwall.com", ch)
	go responseSize("https://stackoverflow.com", ch)
	time.Sleep(3000 * time.Millisecond)
}
