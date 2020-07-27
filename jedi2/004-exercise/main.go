package main

import "fmt"

func main() {
	var x int = 2
	fmt.Printf("%v\t%b\t%#x", x, x, x)
	x = x << 1
	fmt.Println()
	fmt.Printf("%v\t%b\t%#x", x, x, x)

}
