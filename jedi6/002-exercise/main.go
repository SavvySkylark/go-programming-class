package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(ii...))
	fmt.Println(bar(ii))
}

func foo(x ...int) int {
	sum := 0
	for _, num := range x {
		sum += num
	}
	return sum
}

func bar(xs []int) int {
	sum := 0
	for _, num := range xs {
		sum += num
	}
	return sum
}
