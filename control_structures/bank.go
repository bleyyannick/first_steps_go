package main

import "fmt"

func main() {

	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do ?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Println("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is ", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Println("Your deposit: ")
		fmt.Scan(&depositAmount)

		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount: ", accountBalance)
	} else if choice == 3 {
		var moneyOut float64
		fmt.Println(("Withdaw some money: "))
		fmt.Scan(&moneyOut)
		accountBalance -= moneyOut
		fmt.Println("Balance updated! New amount: ", accountBalance)
	}

	fmt.Println("Your choice:", choice)
}
