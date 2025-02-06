package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Width * t.Height) / 2
}
func (t Triangle) Perimeter() float64 {
	return t.Width * 3
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeProperties(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())

}

func main() {

	triangle := Triangle{10, 20}
	var shape Shape = triangle
	PrintShapeProperties(shape)

	rectangle := Rectangle{10, 20}
	var shape_2 Shape = rectangle
	PrintShapeProperties(shape_2)

	circle := Circle{10}
	var shape_3 Shape = circle
	PrintShapeProperties(shape_3)
	/*
		C := Circle{10}
		R := Rectangle{10, 20}
		fmt.Println("Circle:\n")
		PrintShapeProperties(C)
		fmt.Println("-----------------------")
		fmt.Printf("Rectangle: \n")
		PrintShapeProperties(R)
		fmt.Println("-----------------------")
		fmt.Printf("Triangle: \n")
		PrintShapeProperties(Triangle{10, 20})

	*/
}
