package main

import (
	"fmt"
	"strconv"
)

type Element interface{}

type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("List[%d] is an int and its value is %v\n", index, value)
		case string:
			fmt.Printf("List[%d] is an string and its value is %v\n", index, value)
		case Person:
			fmt.Printf("List[%d] is an Person and its value is %v\n", index, value)
		default:
			fmt.Printf("List[%d] is a different type", index)
		}
	}
}
