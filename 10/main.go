package main

import (
	"fmt"

	filemanager "example.com/calculator/fileManager"
	"example.com/calculator/price"
)

func main() {
	taxes := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxes {
		fm := filemanager.New("Pricess.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdmanager := cmdmanager.New()
		priceJob := price.NewPriceIncludeTaxJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process the job")
			fmt.Println(err)
			break
		}
	}
}
