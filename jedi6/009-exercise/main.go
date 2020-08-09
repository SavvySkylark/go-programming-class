package main

import "fmt"

func main() {
	foo(func() {
		fmt.Println("thing")
	})
}

func foo(f func()) {
	f()
}
