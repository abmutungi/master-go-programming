package main

import "fmt"


type person struct {
	name   string
	age    int
	colors []string
	gr     grades
}

type grades struct {
	grade  int
	course string
}

func main() {
	
	me := person{
		name:   "Arnold",
		age:    29,
		colors: []string{"red", "yellow"},
		gr:     grades{grade: 85, course: "Python"},
	}
	you := person{
		name:   "Zakiya",
		age:    26,
		colors: []string{"pink", "blue"},
		gr:     grades{grade: 100, course: "Chemistry"},
	}


	me.gr.grade = 98
	me.gr.course = "Golang"


	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

}