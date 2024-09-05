package ch4

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("ch4/balance.txt", []byte(balanceText), 0644)
}

func readBalanceFromFile() {
	dataRead, err := os.ReadFile("ch4/balance.txt")
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		str_dataRead := string(dataRead)
		float_data, _ := strconv.ParseFloat(str_dataRead, 64)
		fmt.Println(float_data)
	}
}

func throwError() (int, error) {
	return 1, errors.New("custom_error")
}

func Files() {
	writeBalanceToFile(24.5)
	readBalanceFromFile()
	res, err := throwError()
	if err != nil {
		fmt.Printf("-----ERROR----- %v", res)
		// return
		panic("PANIC")
	}
}
