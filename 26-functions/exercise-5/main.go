/*
Change the function from the previous exercise and use a `naked return`.
*/
package main

import "fmt"

func sum(a ...int) (b int) {
	sum := 0

	for _, v := range a {
		sum += v
	}

	return
}

func main() {
	fmt.Println(sum(4, 6, 6))
}
