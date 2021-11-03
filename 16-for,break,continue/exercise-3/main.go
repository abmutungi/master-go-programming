package main

import "fmt"

func main() {
	for i := 1; i < 500; i++ {
		if i%5 == 0 && i%7 == 0 {
			fmt.Println(i)
		}
	}
}
