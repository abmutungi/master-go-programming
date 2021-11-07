package main

import "fmt"

func main() {

	s := "你好 Go!"

	r := []rune(s)

	fmt.Printf("%#v\n", r)

	for i, v := range r {
		fmt.Printf("index: %d -> rune: %c\n", i, v)
	}
}
