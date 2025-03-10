package calculations

import "math"

type Rectangle struct {
	width  float64
	length float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.length)
}

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return t.base / 2 * t.height
}
