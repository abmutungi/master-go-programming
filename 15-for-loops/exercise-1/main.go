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

//Change the code from the previous exercise and use the continue statement to print out all the numbers divisible by 7 between 1 and 50.
// 	for i := 1; i <= 50; i++ {
// 		if i%7 != 0 { // if i is not divisible by 7
// 			continue
// 		}
// 		fmt.Printf("%d ", i)

// 	}
// 	fmt.Println("")

// }
