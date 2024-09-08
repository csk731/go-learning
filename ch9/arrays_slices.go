package ch9

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}
func (p Product) printProd(){
	fmt.Println("Printing Product")
}

func ArraysSlices() {
	prices := [4]float64{1, 2, 3, 4}

	product := Product{
		id:    "1",
		title: "iPhone",
		price: 999.99,
	}
	products := [4]Product{product, product}

	var productNames [2]string
	productNames[0] = "Pixel"
	productNames[1] = "Samsung"

	// var humanNames [2]string
	// humanNames = [2]string{"Chay", "Ram"}
	var humanNames [2]string = [2]string{"Chay", "Ram"}

	fmt.Println("Prices:", prices)
	fmt.Println("Products", products)
	fmt.Println("Product Names:", productNames)
	fmt.Println("Human Name at index 0:", humanNames[0])

	fmt.Println("Length of Prices:",len(prices))

	fmt.Println("Capacity:",cap(prices[:3][:2]),"& Length:",len(prices[:3][:2]))
	fmt.Println("Capacity:",cap(prices[3:]),"& Length:",len(prices[3:]))
	fmt.Println("Capacity:",cap(prices[:3]),"& Length:",len(prices[1:][:1]))
	fmt.Println("Weird Slices:",prices[1:][:1][:3]," & Cap:", cap(prices[1:][:1][:3])) // Capacity will be changed when we slice the array using left part of the colon.

	// Slices
	fmt.Println(prices[0:4])
	fmt.Println(prices[3:])
	fmt.Println(prices[:3])

	// Dynamic Lists
	ages := []float64{1,2,3}
	ages[0]=0
	fmt.Println(ages, &ages[0])
	// ages[3]=2 // Error
	// ages[10]=10 // Error
	ages = append(ages, 5, 6)
	newAges := []float64{44,432,434242,21}
	ages = append(ages,newAges...)
	fmt.Println(ages, &ages[0])

	
}