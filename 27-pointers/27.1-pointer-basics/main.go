package main

import "fmt"

func main() {
	// the & operator aka "address of operator" returns the memory address of a variable

	name := "Zakiya-Rose"
	fmt.Println(&name) // -> 0x14000110210

	//** DECLARING AND INITIALISING POINTERS **//

	var x int = 2
	// the expression &x means the address of x and creates a pointer to an integer variable
	// ptr is of type *int, which is pronounced "pointer to int".

	ptr := &x
	fmt.Printf("ptr is of type %T with value %v and address &%p\n", ptr, ptr, ptr) // -> ptr is of type *int with value 0x14000122008 and address &0x14000122008
	
	
	// declaring a pointer without initialiseing it
	// its zero value is nil

	var ptr1 *float64
	fmt.Println(ptr1)	

	//creating a pointer using new() built-in function
	p := new(int) // ir creates a pointer called p that is a pointer to an int type

	x = 100
	p = &x // intialising p

	fmt.Printf("p is of type %T with value %v\n", p, p) // -> p is of type *int with a value 0x1400012c008

	fmt.Printf("address of x is %p\n", &x) // -> address of x is 0x14000122008

	    //** THE DEREFERENCING OPERATOR **//

		// * in front of a pointer is called the dereferencing operator
		*p = 90 // equivalent to x = 90 because *p = x
		//x and *p is the same thing

		fmt.Println(*p) // => 90
		fmt.Println("*p == x:", *p == x) // => *p == x: true

		fmt.Println("Value of x:", *p) // Value of x: 90, equivalent to fmt.Println(x)

		*p = 10 		 // equivalent to x = 10
		*p = *p /2		 // dividing x through the pointer
		fmt.Println(x) 	 // -> 5

    // In a nutshell:
    // &value => pointer -> if you have a value you turn it into an address or pointer by using the ampersand operator
    // *pointer => value  -> and if you have pointer you turn it into value value by using the star operator
}
