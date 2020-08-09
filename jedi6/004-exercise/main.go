package main

import "fmt"

type person struct {
	first string
	last  string
	age   string
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Skylar",
		last:  "Smith",
		age:   "29",
	}
	p.speak()
}
