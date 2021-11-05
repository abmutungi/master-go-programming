package main

import "fmt"

func main() {
	nums := []float64{1.2, 3.4, 5.6}

	nums = append(nums, 10.1)

	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println(nums)

	n := []float64{23.6, 24.8}

	nums = append(nums, n...)

	fmt.Println(nums)
}
