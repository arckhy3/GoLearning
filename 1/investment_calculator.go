package main

import (
	"fmt"
	"math"
)

const invlationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnReate float64

	fmt.Print("Investment Amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate : ")
	fmt.Scan(&expectedReturnReate)

	fmt.Print("Years : ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculationFutureValues(investmentAmount, expectedReturnReate, years)

	formatFv := fmt.Sprintf("Future value : %.2f\n", futureValue)
	formatFrv := fmt.Sprintf("Future value (adjust for invlation) : %.2f\n", futureRealValue)

	fmt.Print(formatFv, formatFrv)
}

func calculationFutureValues(investmentAmount, expectedReturnReate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnReate/100, years)
	rfv = fv / math.Pow(1+invlationRate/100, years)
	return
}
