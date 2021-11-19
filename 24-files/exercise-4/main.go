package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// opening the file in read-only mode. The file must exist (in the current working directory)
	// use a valid path!
	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}

	// defer closing the file
	defer file.Close()

	// ioutil.ReadAll() reads from the file until an error or EOF and returns the data it read
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
}
