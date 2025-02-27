package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u User) OutputUserDetailMethod() {
	fmt.Println(u.firstName, u.lastName, u.birthdate) //, u.createdAt)
}

func (u *User) ClearUserDataMethod() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
	u.createdAt = time.Time{}
}

func New(firsName, lastName, birthdate string) (*User, error) {
	if firsName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, Last name and Birthdate are required")
	}

	return &User{
		firstName: firsName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
