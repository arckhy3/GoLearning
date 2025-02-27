package list

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	fmt.Println(prices)
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	fmt.Println(prices[3])

	featurePrices := prices[1:3]

	fmt.Println(featurePrices)

	dynamicSlices()
}

func dynamicSlices() {
	prices := []float64{10.99, 8.99}

	fmt.Println(prices)

	prices[1] = 9.99
	fmt.Println(prices)

	prices = append(prices, 12.99)
	fmt.Println(prices)

	discountPrices := []float64{5.99, 6.99, 8.99}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
