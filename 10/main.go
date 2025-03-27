package main

import (
	cmdmanager "example.com/calculator/cmdManager"
	"example.com/calculator/price"
)

func main() {
	taxes := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxes {
		// fm := filemanager.New("Price.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdmanager := cmdmanager.New()
		priceJob := price.NewPriceIncludeTaxJob(cmdmanager, taxRate)
		priceJob.Process()
	}
}
