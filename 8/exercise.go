package main

import "fmt"

type product struct {
	title  string
	id     int
	prices float64
}

func main() {

	//1
	fmt.Println("1")
	hobbies := [3]string{"gaming", "travel", "read novel"}
	fmt.Println(hobbies)
	fmt.Println("------------")

	//2
	fmt.Println("2")
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	fmt.Println("------------")

	//3
	fmt.Println("3")
	//slicesHobies := hobbies[0:2]
	slicesHobies := hobbies[:2]
	fmt.Println(slicesHobies)
	fmt.Println("------------")

	//4
	fmt.Println("4")
	slicesHobies2 := slicesHobies[1:3]
	fmt.Println(slicesHobies2)
	fmt.Println("------------")

	//5
	fmt.Println("5")
	goals := []string{"be better be dev", "learn go"}
	fmt.Println(goals)
	fmt.Println("------------")

	//6
	fmt.Println("6")
	goals[1] = "learn more"
	goals = append(goals, "learn go to dev be")
	fmt.Println(goals)
	fmt.Println("------------")

	//7
	fmt.Println("7")
	products := []product{{title: "product 1", id: 1, prices: 10.99}, {title: "product 2", id: 2, prices: 15.99}}
	fmt.Println(products)
	products = append(products, product{title: "product 3", id: 3, prices: 4.99})
	fmt.Println(products)
	fmt.Println("------------")
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
