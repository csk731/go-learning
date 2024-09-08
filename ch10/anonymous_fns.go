package ch10

import "fmt"

func AnonymousFunctions() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	// transformNumbers(&numbers, func(n int) int {
	// 	return n * 2
	// })
	transformNumbers(&numbers, createTransformerFunction(3))
	fmt.Println(numbers)
}

// Closures - A function that has non local variables is considered as a closure
func createTransformerFunction(multiplier int) transformFn{
	return func (n int) int  {
		return n * multiplier;
	}
}
