package main

import "fmt"

func main() {

	//Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.

	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}

