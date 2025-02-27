package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	//  var appUser user

	// appUser := user.User{
	// 	FirstName: userFirstName,
	// 	LastName:  userLastName,
	// 	Birthdate: userBirthdate,
	// }

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin, err := user.NewAdmin("Admin", "-", "-", "admin@email.com", "admin")

	if err != nil {
		fmt.Println(err)
		return
	}

	admin.OutputUserDetailMethod()
	admin.ClearUserDataMethod()
	admin.OutputUserDetailMethod()

	// ... do something awesome with that gathered data!

	// fmt.Println(firstName, lastName, birthdate)
	// outputUserDetail(appUser)
	appUser.OutputUserDetailMethod()
	appUser.ClearUserDataMethod()
	appUser.OutputUserDetailMethod()
}

// func outputUserDetail(u *user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
