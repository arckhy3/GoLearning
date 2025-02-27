package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers2 := []int{5, 8, 9, 3}
	double := transformValue(&numbers, double)
	triple := transformValue(&numbers, triple)

	fmt.Println(double)
	fmt.Println(triple)

	funcNumber1 := transformFunc(&numbers)
	funcNumber2 := transformFunc(&numbers2)

	newVal1 := transformValue(&numbers, funcNumber1)
	newVal2 := transformValue(&numbers2, funcNumber2)

	fmt.Println(newVal1)
	fmt.Println(newVal2)

	anomFn := transformValue(&numbers2, func(i int) int { return i * 2 })

	fmt.Println("anomimous function: ", anomFn)

	newDouble := newTransform(2)
	newTriple := newTransform(3)

	fmt.Println(transformValue(&numbers, newDouble))
	fmt.Println(transformValue(&numbers, newTriple))

	recursiveFactor := factorial(7)

	fmt.Println(recursiveFactor)

	fmt.Println(sumup(1, 2, 3, 4, 5))
	fmt.Println(sumup(numbers...))

}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func newTransform(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func transformFunc(number *[]int) transformFn {
	if (*number)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func transformValue(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}
	for _, value := range *numbers {
		tNumbers = append(tNumbers, transform(value))
	}

	return tNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
