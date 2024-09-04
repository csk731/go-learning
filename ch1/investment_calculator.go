package ch1

import (
	"fmt"
	"math"
)

func InvestmentCalculator() /*float64*/ {
	const inflationRate = 2.5
	// var investmentAmount, years float64 = 1000, 10
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	// var years float64 = 10

	fmt.Print("Enter Inverstment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Years: ")
	fmt.Scan(&years)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	// return futureRealValue
}

/**

References:
1. https://dev.to/diwakarkashyap/packages-and-import-in-go-golang-3adk#:~:text=In%20Go%2C%20you%20can%20use,packages%20in%20your%20source%20code.

**/
