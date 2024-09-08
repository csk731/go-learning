package ch8

import (
	"fmt"
	"math"
)

/*
Shape - interface
*/
type Shape interface {
	GetArea() float64
}

type Shape2 interface {
	Shape
	GetPerimeter() float64
}

/*
Circle - struct
*/
type Circle struct {
	Radius float64
}

func (c Circle) GetArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

/*
Rectangle - interface
*/
type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) GetArea() float64 {
	return r.Length * r.Breadth
}

func GetTotalArea(s Shape) float64 {
	// switch s.(type) {
	// case Rectangle:
	// 	fmt.Println("This is Rectangle")
	// case Circle:
	// 	fmt.Println("This is Circle")
	// default:
	// 	fmt.Println("Nothing Matched.!")
	// }
	typedValue, ok := s.(Rectangle)
	fmt.Println(typedValue, ok)
	return s.GetArea()
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func Interfaces() {
	realCircle := Circle{
		Radius: 5,
	}
	realRectangle := Rectangle{
		Length:  5,
		Breadth: 10,
	}
	res1 := GetTotalArea(realCircle)
	fmt.Println(res1)
	res2 := GetTotalArea(realRectangle)
	fmt.Println(res2)

	fmt.Println(add(1, 2))
} 
