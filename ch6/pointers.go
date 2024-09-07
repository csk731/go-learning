package ch6

import "fmt"

func Pointers() {
	// var x int
	// y, _ := fmt.Scan(&x)
	// fmt.Println(x, y)

	age := 32
	// agePointer := &age
	var agePointer *int = &age
	fmt.Println("Age Pointer:", agePointer)
	fmt.Println("Age:", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println("Age Pointer:", agePointer)
	fmt.Println("Age:", *agePointer)

	// fmt.Println("Age:",age,&age)
	// getAdultYears(age)
	// fmt.Println("Age:",age,&age)

}

func getAdultYears(age int) int {
	fmt.Println("Age:", age, &age)
	age = age - 18
	return age
}

func editAgeToAdultYears(agePointer *int) {
	*agePointer = *agePointer - 18
}
