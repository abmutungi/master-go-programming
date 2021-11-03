package main

import "fmt"

func main() {
	// Using the var keyword, declare an array called cities
	// with 2 elements of type string. Don't initialize the array.
	// Print out the cities array and notice the value of its elements

	var cities [2]string
	fmt.Printf("%#v\n", cities)

	// Declare an array called grades of type [3]float64 and initialize only
	// the first 2 elements using an array literal.
	// Print out the cities array and notice the value of its elements

	grades := [3]float64{4.5, 9.7}
	fmt.Printf("%#v\n", grades)

	// Declare an array called taskDone
	// using the ellipsis operator. The elements are of type bool.

	taskDone := [...]bool{true, false}
	fmt.Printf("%#v\n", taskDone)

	// Iterate over the array called cities using the classical
	// for loop syntax and len function and print out element by element.

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	// Iterate over grades using the range
	// keyword and print out element by element.

	for index, value := range grades {
		fmt.Println(index, value)
	}
}
