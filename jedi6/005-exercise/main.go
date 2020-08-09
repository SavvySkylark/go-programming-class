package main

import (
	"fmt"
	"math"
)

type square struct {
	l float64
	w float64
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (s square) area() float64 {
	return s.l * s.w
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		l: 1.0,
		w: 1.0,
	}

	c := circle{
		r: 2,
	}

	info(s)
	info(c)
}
