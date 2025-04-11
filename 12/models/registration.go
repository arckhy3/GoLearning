package models

import (
	"example.com/event/db"
)

type Register struct {
	ID      int64
	EventID int64
	UserID  int64
}

func (register *Register) Save() error {
	query := `
	INSERT INTO registrations (event_id,user_id)
	VALUES(?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(register.EventID, register.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	register.ID = id
	return nil
}

func (register Register) Delete() error {
	query := `
	DELETE FROM registrations
	WHERE event_id = ?
	AND
	user_id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(register.EventID, register.UserID)

	return err
}
