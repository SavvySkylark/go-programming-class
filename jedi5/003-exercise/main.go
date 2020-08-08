package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIcecreamFlavors []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(t.doors)
	fmt.Println(t.color)
	fmt.Println(t.fourWheel)

	fmt.Println(s.color)
	fmt.Println(s.doors)
	fmt.Println(s.luxury)
}
