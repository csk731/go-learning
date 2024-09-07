package ch7

import "fmt"

type customString string // Where we use? Which is an extra overhead. - But when we want to attach methods to built in types.

func (text customString) log() {
	fmt.Println(text)
}

func Aliases() {
	var name customString = "Chay"
	fmt.Println(name)
	name.log()
}
