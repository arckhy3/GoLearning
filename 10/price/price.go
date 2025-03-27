package price

import (
	"errors"
	"fmt"

	iomanager "example.com/calculator/IOManager"
	"example.com/calculator/conversion"
)

type PriceIncludeTaxJob struct {
	IOManager       iomanager.IOManager `json:"-"`
	InputPrice      []float64           `json:"input_price"`
	TaxRate         float64             `json:"tax_rate"`
	PriceIncludeTax map[string]string   `json:"price_include_tax"`
}

func NewPriceIncludeTaxJob(iom iomanager.IOManager, taxRate float64) *PriceIncludeTaxJob {
	return &PriceIncludeTaxJob{
		IOManager:  iom,
		InputPrice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}
}

func (job *PriceIncludeTaxJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrice {
		priceIncludeTax := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", priceIncludeTax)
	}

	job.PriceIncludeTax = result
	return job.IOManager.WriteData(job)
}

func (job *PriceIncludeTaxJob) LoadData() error {

	lines, err := job.IOManager.LoadData()

	if err != nil {
		err = errors.New("Failed to read data")
		return err
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		err = errors.New("Value failed to convert")
		return err
	}

	job.InputPrice = prices
	return nil
}
