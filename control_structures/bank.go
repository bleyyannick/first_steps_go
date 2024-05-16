package main

import (
	"fmt"
	"os"
	"strconv"
)

func writeBalanceToAFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.text", []byte(balanceText), 0644)
}

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile("balance.txt")
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance

}

func main() {

	accountBalance := getBalanceFromFile()
	for i := 0; i < 2; i++ {
		fmt.Println("Welcome to Go Bank")
		fmt.Println("What do you want to do ?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToAFile(accountBalance)
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
			writeBalanceToAFile(accountBalance)
		default:
			fmt.Println("Goodbye !")
			return
		}
		fmt.Println("Your choice:", choice)
	}

}
