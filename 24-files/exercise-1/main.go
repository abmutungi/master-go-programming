package main

import (
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("aa.txt")
	// error handling
	if err != nil {
		// log the error and exit the program
		log.Fatal(err) // the idiomatic way to handle errors
	}
	newFile.Close()
}
