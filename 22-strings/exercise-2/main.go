package main

import "fmt"

func main() {
	r := 'ă'
	fmt.Printf("%T\n", r)

	s, t := "ma", "m" // declaring 2 strings

	fmt.Println(t + s + string(r))
}
