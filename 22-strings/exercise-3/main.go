package main

import "fmt"

func main() {

	s1 := "țară means country in Romanian"

	// iterating over the string to print out byte by byte
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v ", s1[i])
	}

	fmt.Println()

	// iterating over the string and print out rune by rune
	// and the byte index where the rune starts in the string
	for i, v := range s1 {
		fmt.Printf("byte index: %#d -> rune: %c\n", i, v)
	}
	fmt.Println()

}
