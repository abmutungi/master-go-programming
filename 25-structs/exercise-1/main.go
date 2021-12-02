package main

import "fmt"

func main() {
	type person struct {
		name    string
		age     int
		colours []string
	}

	me := person{name: "Arnold", age: 29, colours: []string{"red", "black"}}
	you := person{name: "Zakiya", age: 26, colours: []string{"pink", "blue"}}

	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)
}
