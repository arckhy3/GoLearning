package main

import (
	"example.com/calculator/price"
)

func main() {
	taxes := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxes {
		priceJob := price.NewPriceIncludeTaxJob(taxRate)
		priceJob.Process()
	}
}
