/*
Change the function from the previous exercise to do a case-insensitive search.
*/
package main

import (
	"fmt"
	"strings"
)

func searchItem(a []string, b string) bool {
	for _, v := range a {
		if strings.EqualFold(v, b) {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"Lion", "tiger", "bear"}
	result := searchItem(animals, "beaR")

	fmt.Println(result)
}

