package price

import (
	"fmt"

	"example.com/calculator/conversion"
	filemanager "example.com/calculator/fileManager"
)

type PriceIncludeTaxJob struct {
	IOManager       filemanager.FileManager `json:"-"`
	InputPrice      []float64               `json:"input_price"`
	TaxRate         float64                 `json:"tax_rate"`
	PriceIncludeTax map[string]string       `json:"price_include_tax"`
}

func NewPriceIncludeTaxJob(fm filemanager.FileManager, taxRate float64) *PriceIncludeTaxJob {
	return &PriceIncludeTaxJob{
		IOManager:  fm,
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
	job.IOManager.WriteData(job)
}

func (job *PriceIncludeTaxJob) LoadData() {

	lines, err := job.IOManager.LoadData()

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
