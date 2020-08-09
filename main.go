package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7}

	x := odd(sum, xi...)

	fmt.Println(x)
}

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func odd(f func(xi ...int) int, vi ...int) int {
	yi := []int{}

	for _, i := range vi {
		if i%2 == 1 {
			yi = append(yi, i)
		}
	}
	return f(yi...)
}
