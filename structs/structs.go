package main

import (
	"fmt"
	"time"
)

// Data type struct
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u user) outputDetailsUser() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearDetailsUser() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthDate,
		createdAt: time.Now(),
	}

	appUser.outputDetailsUser()
	appUser.clearDetailsUser()
	appUser.outputDetailsUser()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
