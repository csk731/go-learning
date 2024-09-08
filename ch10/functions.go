package ch10

import (
	"fmt"
)

type transformFn func(int) int

func Functions() {
	numbers := []int{1, 2, 3, 4, 5}
	// doubleNumbers(&numbers)
	// transformNumbers(&numbers, double)
	transformNumbers(&numbers, triple)
	fmt.Println(numbers)

	// ans := addAll(1,2,3,4,5)
	ans := addAll(numbers...)
	fmt.Println(ans)

}

func doubleNumbers(numbers *[]int) {
	// fmt.Println(&(*numbers)[0])
	for idx, val := range *numbers {
		(*numbers)[idx] = double(val)
	}
	// fmt.Println(&(*numbers)[0])
}

func double(i int) int {
	return i * 2
}

func triple(i int) int {
	return i * 3
}

func transformNumbers(numbers *[]int, transform transformFn) {
	for idx, val := range *numbers {
		(*numbers)[idx] = transform(val)
	}
}

func getTransformerFunction() transformFn {
	return double
}

// Variadic functions
func addAll(nums ...int) int {
	res := 0
	for _, val := range nums{
		res += val;
	}
	return res
}
