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

	p2 := person{
		firstName: "Greyson",
		lastName:  "Smith2",
		favoriteIcecreamFlavors: []string{
			"Chocolate",
			"Vanillia",
			"Strawberry",
		},
	}
	m := map[string]person{
		p.lastName:  p,
		p2.lastName: p2,
	}

	for _, person := range m {
		fmt.Println(person.firstName)
		fmt.Println(person.lastName)
		for _, flavor := range person.favoriteIcecreamFlavors {
			fmt.Println(flavor)
		}
	}
}
