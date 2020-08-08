package main

import "fmt"

func main() {
	x := struct {
		field1 int
		field2 string
	}{
		field1: 1,
		field2: "thing",
	}

	fmt.Println(x.field1)
	fmt.Println(x.field2)
}
