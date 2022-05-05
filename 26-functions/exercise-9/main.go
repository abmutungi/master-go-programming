/*Create a function that takes in an int value and prints out that value.

Assign the function to a variable, print out the type of the variable
and then call that function through the variable name.
*/
package main

import "fmt"

func ok(a int) {
	fmt.Println(a)
}

func main() {
	v := ok

	fmt.Printf("%T\n", v)

	v(9)
}
