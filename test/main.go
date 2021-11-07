package main

import "fmt"

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	count := 0

	var input string
	fmt.Scanln(&input)

	results = append(results, input)

	for _, v := range results {
		if v == "w" {
			count += 3
		} else if v == "d" {
			count++
		} else if v == "l" {
			count += 0
		}
	}
	fmt.Println(count)
}
