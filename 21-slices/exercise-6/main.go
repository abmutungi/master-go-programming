package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := make([]string, len(friends))
	copy(yourFriends, friends)

	yourFriends[0] = "Zakiya" // the slices are not connected by modifying one slice

	fmt.Println(friends, yourFriends)
}
