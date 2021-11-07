package main

import "fmt"

func main() {
	i := 3
	f := 3.2

	// int to float64
	f1 := float64(i)
	fmt.Printf("f1's type: %T, f1's value: %f\n", f1, f1)

	// float64 to int
	i1 := int(f)
	fmt.Printf("i1's type: %T, i1's value: %d\n", i1, i1)
}
