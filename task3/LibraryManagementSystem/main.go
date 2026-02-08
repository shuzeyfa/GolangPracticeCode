package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) printname() (string, int) {
	return p.name, p.age
}

func main() {

	first := Person{
		name: "first",
		age:  12,
	}

	fmt.Println(first.name, first.age)
	fmt.Println(first.printname())

}
