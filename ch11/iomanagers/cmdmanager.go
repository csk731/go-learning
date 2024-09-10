package iomanagers

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println(`
			Please enter your prices. 
			Confirm every price with ENTER. 
			Once you are done. Enter "-1" and OK`)

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "-1" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func NewCMDManager() CMDManager {
	return CMDManager{}
}
