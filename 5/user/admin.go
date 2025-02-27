package user

import "errors"

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(firsName, lastName, birthdate, email, password string) (*Admin, error) {
	if firsName == "" || lastName == "" || birthdate == "" || email == "" || password == "" {
		return nil, errors.New("Admin Data Required")
	}

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: firsName,
			lastName:  lastName,
			birthdate: birthdate,
		},
	}, nil
}
