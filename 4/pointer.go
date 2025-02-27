package main

import "fmt"

func main() {
	age := 34

	// var agePointer *int

	agePointer := &age

	fmt.Println("Age: ", *agePointer)

	// adultYear := getAdultYear(agePointer)
	getAdultYear(agePointer)
	fmt.Println("Years turn adult: ", age)
}

func getAdultYear(age *int) {
	// return (*age - 18)
	*age -= 18
}
