package main

import "fmt"

func presentOption() {
	optionMenu("What do you want to do?")
	optionMenu("1. Check Balance")
	optionMenu("2. Deposit Money")
	optionMenu("3. Withdraw Money")
	optionMenu("4. Exit")
}

func optionMenu(menu string) {
	fmt.Println(menu)
}
