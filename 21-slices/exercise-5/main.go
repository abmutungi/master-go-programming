package main

import "fmt"

func main() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	slice := nums[1:5]

	sum := 0

	for _, v := range slice {
		fmt.Println(v)
		sum += v

	}
	fmt.Println(sum)
}
