package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	sq := square{2.1}
	ci := circle{2.1}

	printArea(sq)
	printArea(ci)
}
