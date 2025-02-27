package main

import (
	"fmt"

	filemanager "example.com/calculator/fileManager"
	"example.com/calculator/price"
)

func main() {
	taxes := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxes {
		fm := filemanager.New("Price.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := price.NewPriceIncludeTaxJob(fm, taxRate)
		priceJob.Process()
	}
}
