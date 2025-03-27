package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) LoadData() ([]string, error) {
	fmt.Println("Please enter your price. Confirm every price with enter")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteData(data any) error {
	fmt.Println(data)

	return nil
}

func New() CMDManager {
	return CMDManager{}
}
