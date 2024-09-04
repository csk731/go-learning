package ch1

import "fmt"

func ProfitCalculator() {
	var revenue float64
	var expenses float64
	var tax_rate float64

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&tax_rate)

	earnings_before_tax := revenue - expenses
	profit := earnings_before_tax * (1-tax_rate)
	ratio := earnings_before_tax / profit

	fmt.Println("Earning Before Tax: ", earnings_before_tax)
	fmt.Printf("Profit: %1.2v\n", profit)
	fmt.Print("Ratio:", ratio)
}
