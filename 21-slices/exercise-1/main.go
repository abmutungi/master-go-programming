package main

import "fmt"

func main() {
	mutungi := []string{"Arnold", "Natasha", "Alexandra"}

	for i, v := range mutungi {
		fmt.Printf("Index: %d, Child: %q\n", i, v)
	}
}
