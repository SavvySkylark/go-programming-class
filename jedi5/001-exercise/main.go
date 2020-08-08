package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIcecreamFlavors []string
}

func main() {
	p := person{
		firstName: "Skylar",
		lastName:  "Smith",
		favoriteIcecreamFlavors: []string{
			"Chocolate",
			"Vanillia",
			"Strawberry",
		},
	}

	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
	for _, iceCream := range p.favoriteIcecreamFlavors {
		fmt.Println(iceCream)
	}
}
