package main

import (
	"fmt"
)

type Shape interface {
	GetArea() float64
	GetPerimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) GetArea() float64{
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) GetPerimeter() float64{
	return 2 * 3.14 * c.Radius
}
// Since Circle has the methods GetArea and GetPerimeter with the correct signatures, it implicitly implements the Shape interface.
// So we can say that a circle is a shape.

type Square struct {
	Side float64
}
func (s *Square) GetArea() float64{
	s.Side = 10 // Modifying the Side field to demonstrate pointer receiver
	return s.Side * s.Side
}
func (s Square) GetPerimeter() float64{
	return 4 * s.Side
}

func dummyFunction(s *Square){
	s.Side = 20
}

func NewSquare()Square{
	return Square{
		Side:5,
	}
}
func main(){
	square := NewSquare()
	dummyFunction(&square)
	// square.GetArea()
	
	// PrintShapeArea(square) // send the address of square, because getArea has a pointer receiver
	// fmt.Println("Side of the square after GetArea call: ", square.Side)
}

// func PrintShapeArea(s Shape){
// 	fmt.Println("Area: ", s.GetArea())
// }




