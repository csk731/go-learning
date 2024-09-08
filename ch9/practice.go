package ch9

import (
	"fmt"
)

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

func Practice() {
	// T1
	var hobbies [3]string = [3]string{"Code", "Music", "Badminton"}
	fmt.Println(hobbies, &hobbies[1], &hobbies[2])

	// T2
	fmt.Println(hobbies[0])
	var fewHobbies []string = hobbies[1:3]
	fmt.Println(fewHobbies, &fewHobbies[0], &fewHobbies[1])

	// T3
	hobbiesSlice1 := hobbies[0:2]
	fmt.Println(hobbiesSlice1, &hobbiesSlice1[0], cap(hobbiesSlice1))
	hobbiesSlice2 := hobbies[:2]
	fmt.Println(hobbiesSlice2, &hobbiesSlice2[0])
	hobbiesSlice3 := make([]string, 2)
	copy(hobbiesSlice3, hobbiesSlice2)
	fmt.Println(hobbiesSlice3, &hobbiesSlice3[0])

	// T4
	reSlice := hobbiesSlice1[1:3]
	fmt.Println(reSlice)

	// T5
	goals := []string{"MAANG","SDE"}
	fmt.Println(goals)
	goals[1] = "SE"
	goals = append(goals, "US")
	fmt.Println(goals)

	// T6
	prod1 := Product{
		id: "1",
		title: "ABC",
		price: 12.00,
	}
	prod2 := Product{
		id: "2",
		title: "ABCD",
		price: 10.00,
	}
	products := []Product{prod1, prod2}
	fmt.Println(products)
	prod3 := Product{
		id: "3",
		title: "ABCDE",
		price: 5.00,
	}
	products = append(products, prod3)
	fmt.Println(products)
}
