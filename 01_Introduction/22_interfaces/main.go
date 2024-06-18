package main

import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (r square) perimeter() float64 {
	return r.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.perimeter())
}

func main() {
	shapes := []shape{
		square{length: 10.5},
		square{length: 5.5},
		circle{radius: 4},
		circle{radius: 10},
	}

	for _, v := range shapes {

		printShapeInfo(v)
		fmt.Println("---")
	}
}
