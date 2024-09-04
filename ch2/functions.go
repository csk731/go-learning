package ch2

import "fmt"

var x int = 10

func Functions(){
	outputText("Hello, ch2")
	var a, b int = ReturnTwoValues()
	fmt.Println(a, b)
	VarArgs("hi","bye")
	fmt.Println(ReturntwoValues2())
}

func outputText(text string){
	fmt.Println(text);
	fmt.Println(x);
}

func ReturnTwoValues() (int, int){
	return 1,2
}

func VarArgs(x ...string){
	fmt.Println(x)
}

func ReturntwoValues2() (fv int, fv1 int){
	fv = 10
	fv1 =20
	return
}
