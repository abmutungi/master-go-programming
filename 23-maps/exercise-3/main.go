package main

import "fmt"

func main() {

	m := map[int]bool{1: true, 2: false, 3: false}

	delete(m, 2)

	for i, v := range m {
		fmt.Println(i, ":", v)
	}
}
