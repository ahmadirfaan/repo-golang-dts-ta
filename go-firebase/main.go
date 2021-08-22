package main

import (
	"fmt"
	"github.com/ahmadirfaan/go-firebase/firebase"
)


func main() {
	bucket := firebase.CloudStorage()
	fmt.Println(bucket)
}