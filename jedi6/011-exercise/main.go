package main

import "fmt"

func main() {
	fmt.Println(factorial(4, 1))
}

// 4! 4*3*2*1

func factorial(f int, acc int) int {

	if f <= 0 {
		return acc
	}
	acc *= f
	f--
	return factorial(f, acc)

}
