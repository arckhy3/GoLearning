package main

import (
	"fmt"
)

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userName := make([]string, 2, 5)
	userName[0] = "julie"
	userName = append(userName, "john")
	userName = append(userName, "peter")

	fmt.Println(userName)

	// courseRating := map[string]float64{}

	// courseRating := make(map[string]float64, 2)

	courseRating := make(floatMap, 3)

	courseRating["go"] = 4.9
	courseRating["react"] = 4.8
	courseRating["js"] = 4.9

	// fmt.Println(courseRating)

	courseRating.output()

	for i, val := range userName {
		fmt.Println("index: ", i)
		fmt.Println("value: ", val)
	}

	for key, val := range courseRating {
		fmt.Println("key: ", key)
		fmt.Println("value: ", val)
	}
}
