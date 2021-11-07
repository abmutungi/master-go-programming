package main

import "fmt"

func main() {
	s := "你好 Go!"

	s1 := []byte(s)

	fmt.Printf("%T\n", s1)

	for i, v := range s1 {
		fmt.Printf("index: %d -> byte: %c\n", i, v)
	}
}
