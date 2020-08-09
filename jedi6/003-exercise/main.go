package main

import "fmt"

func main() {
	foo()
}

func foo() {
	defer bar()
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
