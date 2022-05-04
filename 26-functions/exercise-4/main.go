/*
Create a function with the identifier sum that takes in a variadic parameter
of type int and returns the sum of all values of type int passed in.
*/

package main

import "fmt"

func sum(a ...int) int {
	sum := 0

	for _, v := range a {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(sum(4, 6, 6))
}
