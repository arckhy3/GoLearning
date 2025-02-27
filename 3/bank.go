package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFileName = "balance.txt"

func main() {
	var choice int
	accountBalance, err := fileops.ReadFloatFromFile(balanceFileName)

	optionMenu("Welcome to Bank")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for choice < 4 {
		println("--------------")
		if err != nil {
			fmt.Println(err)
			fmt.Println("-------------")
		}

		presentOption()

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance:", accountBalance)
		case 2:
			fmt.Print("Your Deposit:")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Deposit Amount must be greater than 0!")
				continue
			}
			accountBalance += deposit
			fmt.Println("Balance Updated, New Balance:", accountBalance)
			fileops.WriteFLoatToFile(balanceFileName, accountBalance)
		case 3:
			fmt.Print("Withdraw Amount:")
			var amountWithdraw float64
			fmt.Scan(&amountWithdraw)
			if amountWithdraw > accountBalance {
				fmt.Println("Insuficient Amount")
				continue
			}
			if amountWithdraw <= 0 {
				fmt.Println("Deposit Amount must be greater than 0!")
				continue
			}
			accountBalance -= amountWithdraw
			fmt.Println("Balance Update, New Balance:", accountBalance)
			fileops.WriteFLoatToFile(balanceFileName, accountBalance)
		default:
			fmt.Println("Goodbye!!")
			return
		}
	}
}
