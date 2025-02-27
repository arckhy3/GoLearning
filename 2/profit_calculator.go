package main

import (
	"errors"
	"fmt"
	"os"
)

const ebtFilName = "earningbeforetax.txt"
const eatFileName = "earningaftertax.txt"
const profitRatioFileName = "profitratio.txt"
const resultFileName = "result.txt"

func main() {

	revenue, err := populateVariable("Revenue : ")
	if err != nil {
		panic(err)
	}
	expenses, err := populateVariable("Expenses : ")
	if err != nil {
		panic(err)
	}
	taxRate, err := populateVariable("Tax Rate : ")
	if err != nil {
		panic(err)
	}

	ebt, eat, profitRatio := profitCalculation(revenue, expenses, taxRate)

	fmt.Printf("Earning Before Tax : %.2f\n", ebt)
	fmt.Printf("Earning After Tax : %.2f\n", eat)
	fmt.Printf("Ratio : %.2f", profitRatio)

	fmt.Println()
	writeprofitToFile(ebt, eat, profitRatio)
}

func populateVariable(mesage string) (value float64, errMsg error) {
	errMsg = nil
	fmt.Print(mesage)
	fmt.Scan(&value)
	if value <= 0 {
		errMsg = errors.New("invalid value. value must be greater than 0")
	}
	return
}

func profitCalculation(revenue, expenses, taxRate float64) (ebt float64, eat float64, profitRatio float64) {
	ebt = revenue - expenses
	eat = ebt - (1 + taxRate/100)
	profitRatio = ebt / eat
	return
}

func writeprofitToFile(ebt, eat, profitRatio float64) {
	ebtText := fmt.Sprintf("%2.f", ebt)
	eatText := fmt.Sprintf("%.2f", eat)
	profitRatioText := fmt.Sprintf("%.2f", profitRatio)
	result := fmt.Sprintf("EBT : %.2f\nEAT : %.2f\nRATIO : %.2f\n", ebt, eat, profitRatio)
	os.WriteFile(ebtFilName, []byte(ebtText), 0644)
	os.WriteFile(eatFileName, []byte(eatText), 0644)
	os.WriteFile(profitRatioFileName, []byte(profitRatioText), 0644)
	os.WriteFile(resultFileName, []byte(result), 0644)
}
