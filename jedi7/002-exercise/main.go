package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.last = "changed"
	p.first = "changed"
}

func main() {
	p := person{
		first: "Skylar",
		last:  "Smith",
	}
	fmt.Println(p)
	changeMe(&p)

	fmt.Println(p)
}
