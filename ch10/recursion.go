package ch10

import "fmt"

func Recursion(){
	fmt.Println(fact(5))
}

func fact(i int) int{
	if i==0 || i==1 {
		return 1
	}
	return i * fact(i-1)
}