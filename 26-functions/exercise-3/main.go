/*
Write a function called myFunc() that takes exactly one argument which is an int number written between
double quotes (this is in fact a string). If the argument is integer 'n', the function should return the
result of n + nn + nnn

Example: myFunc('5') returns 5 + 55 + 555 which is 615 and myFunc('9') returns 9 + 99 + 999 which is 1107

*/
package main

import (
	"fmt"
	"strconv"
)

func myFunc(n string) int {
	i, _ := strconv.Atoi(n)
	ii, _ := strconv.Atoi(n + n)
	iii, _ := strconv.Atoi(n + n + n)

	return i + ii + iii
}

func main() {
	fmt.Println(myFunc("6"))
}
