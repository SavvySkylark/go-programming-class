package main

import "fmt"

func main() {
	x := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, s := range x {
		for _, str := range s {
			fmt.Println(str)
		}
	}
}
