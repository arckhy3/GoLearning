package price

import (
	"fmt"

	"example.com/calculator/conversion"
	filemanager "example.com/calculator/fileManager"
)

type PriceIncludeTaxJob struct {
	InputPrice      []float64
	TaxRate         float64
	PriceIncludeTax map[string]string
}

func NewPriceIncludeTaxJob(taxRate float64) *PriceIncludeTaxJob {
	return &PriceIncludeTaxJob{
		InputPrice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}
}

func (job *PriceIncludeTaxJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrice {
		priceIncludeTax := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", priceIncludeTax)
	}

	job.PriceIncludeTax = result
	filemanager.WriteData(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func (job *PriceIncludeTaxJob) LoadData() {

	lines, err := filemanager.LoadData("price.txt")

	if err != nil {
		fmt.Println("Failed to read data")
		fmt.Println(err)
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println("Value failed to convert")
		fmt.Println(err)
	}

	job.InputPrice = prices
}
