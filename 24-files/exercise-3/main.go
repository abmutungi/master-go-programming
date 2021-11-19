package main

import (
	"log"
	"os"
)

func main() {
	// removing the file
	err := os.Remove("data.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
