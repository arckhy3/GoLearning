package models

import (
	"errors"

	"example.com/event/db"
	"example.com/event/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `
	INSERT INTO users(email,password)
	VALUES(?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	user.ID = id

	return nil
}

func (user *User) ValidateCredential() error {
	query := `
	SELECT id,password FROM users
	WHERE
	email = ?
	`
	row := db.DB.QueryRow(query, user.Email)

	var retrievePassword string

	err := row.Scan(&user.ID, &retrievePassword)

	if err != nil {
		return err
	}

	valid := utils.CheckPasswordHash(user.Password, retrievePassword)

	if !valid {
		err = errors.New("invalid credential")
	}

	return err
}

func GetAllUser() ([]User, error) {
	var users = []User{}
	query := `SELECT * FROM users`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user = User{}

		err = rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func FindUserByID(id int64) (*User, error) {
	user := User{}
	query := `
	SELECT * FROM users
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (user User) UpdateByID() error {
	query := `
	UPDATE users
	SET email = ?
	password = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.Password, user.ID)

	return err
}

func (user User) Delete() error {
	query := `DELETE FROM users
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	return err
}
