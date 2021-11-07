package main

import (
	"fmt"
)

func main() {

	var name string = "Arnold Mutungi"

	country := "United Kingdom"

	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)

	//equivalent but using ``:
	fmt.Printf(`Your name: %s
Country: %s
`, name, country)

	fmt.Println("He says: \"Hello\"")
	fmt.Println("C:\\Users\\a.txt")

}
