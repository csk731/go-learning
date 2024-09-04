package main

import (
	"fmt"
	ch1 "go-learning/ch1"
	"reflect"
)

func main() {
	fmt.Println("Hello World")
	var x float64 = 10
	fmt.Println(x, "->", reflect.TypeOf(x))

	y := 199.5
	fmt.Println(y, "->", reflect.TypeOf(y))

	var a, b = "abc", 10
	fmt.Println(a, "->", reflect.TypeOf(a))
	fmt.Println(b, "->", reflect.TypeOf(b))

	a1, b1 := "abc", 10
	fmt.Println(a1, "->", reflect.TypeOf(a1))
	fmt.Println(b1, "->", reflect.TypeOf(b1))

	// ch1.InvestmentCalculator()
	ch1.ProfitCalculator()
}
