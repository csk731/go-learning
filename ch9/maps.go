package ch9

import (
	"fmt"
)

type stringMap map[string]string

func (sm stringMap) output() {
	fmt.Println(sm)
}

func Maps() {
	var website_urls map[string]string = map[string]string{
		"google": "www.google.com",
		"amazon": "www.amazon.com",
	}
	fmt.Println(website_urls)
	fmt.Println(website_urls["google"])
	website_urls["linkedin"] = "www.linkedin.com"
	delete(website_urls, "google")
	fmt.Println(website_urls)

	// x := make([]int,0,2)

	// x = append(x,0)
	// fmt.Println(x)
	// fmt.Println(&x[0])

	// x = append(x, 1)
	// fmt.Println(x)
	// fmt.Println(&x[0])

	// x = append(x, 2)
	// fmt.Println(x)
	// fmt.Println(&x[0])

	myMap := make(map[string]string, 2)
	myMap["x"] = "x"
	myMap["y"] = "y"
	myMap["z"] = "z"
	fmt.Println(myMap)

	myStringMap := make(stringMap, 2)
	fmt.Println(myStringMap)
	myStringMap["x"] = "x"
	myStringMap.output()

	prod21 := Product{
		id:    "2",
		title: "ABCD",
		price: 10.00,
	}
	prod21.printProd()

	for key, val := range website_urls {
		fmt.Println(key, val)
	}
}
