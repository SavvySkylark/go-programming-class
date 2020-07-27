package main

import "fmt"

func main() {
	x := 1
	y := 2

	exp1 := x == y
	exp2 := x <= y
	exp3 := x >= y
	exp4 := x != y
	exp5 := x < y
	exp6 := x > y

	fmt.Println(exp1)
	fmt.Println(exp2)
	fmt.Println(exp3)
	fmt.Println(exp4)
	fmt.Println(exp5)
	fmt.Println(exp6)
}
