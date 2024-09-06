package ch6

import "fmt"

func Pointers(){
	var x int
	y, _ := fmt.Scan(&x)
	fmt.Println(x, y)
}