package essentials

import (
	"fmt"
	"math"
)

// https://gobyexample.com/interfaces
type geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func Measure(g geometry) (float64, float64) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
	return g.area(), g.perim()
}
