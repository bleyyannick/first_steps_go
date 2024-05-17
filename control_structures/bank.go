package main

import (
	"control_structures/bank/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Your phone number", randomdata.PhoneNumber())

	for i := 0; i < 2; i++ {

		presentationOption()

		var choice int
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Println("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid number. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToAFile(accountBalance, accountBalanceFile)
		case 3:
			var moneyOut float64
			fmt.Println(("Withdaw some money: "))
			fmt.Scan(&moneyOut)

			if moneyOut <= 0 {
				fmt.Println("Invalid number. Must be greater than 0")
				continue
			}

			if moneyOut >= accountBalance {
				fmt.Println("Invalid withdraw. You can't withdraw more than you have")
				continue
			}
			accountBalance -= moneyOut
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToAFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye !")
			return
		}
		fmt.Println("Your choice:", choice)
	}

}
