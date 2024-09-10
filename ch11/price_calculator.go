package ch11

import (
	"fmt"
	"go-learning/ch11/iomanagers"
	"go-learning/ch11/prices"
)

func PriceCalculator() error {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	for _, taxRate := range taxRates {
		inputFilePath := "ch11/price.txt"
		outputFilePath := fmt.Sprintf("ch11/result_of_%v%%_tax.json", taxRate)
		fm := iomanagers.NewFileManager(inputFilePath, outputFilePath)
		// cmdm := iomanagers.NewCMDManager()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("ERROR: Could not process a job")
			return err
		}
	}
	return nil
}
