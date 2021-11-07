package main

import "fmt"

func main() {
	mySlice := []float64{1.2, 5.6}

	mySlice[1] = 6

	a := float64(10)
	mySlice[0] = a
	fmt.Println(mySlice)

	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)
}

//[10.0 6.0]
//[10.0 10.10 10.0]
