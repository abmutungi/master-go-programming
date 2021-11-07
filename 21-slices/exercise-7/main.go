package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := []string{}

	yourFriends = append(yourFriends, friends...)

	yourFriends[0] = "Dan" // the slices are not connected by modifying one slice

	fmt.Println(friends, yourFriends)
}
