package models

import (
	"time"

	"example.com/event/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e Event) Save() (Event, error) {
	query := `
	INSERT INTO events(name,description,location,datetime,user_id)
	VALUES(?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return e, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return e, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return e, err
	}
	e.ID = id
	return e, err
}

func GetEventByID(id int64) (*Event, error) {
	var event = Event{}
	query := "SELECT * FROM events WHERE ID = ?"

	row := db.DB.QueryRow(query, id)

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetAll() ([]Event, error) {
	var events = []Event{}
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return events, err
	}
	defer rows.Close()
	for rows.Next() {
		var event = Event{}
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, err
}

func (e Event) UpdateByID() error {
	query := `
				UPDATE events
				SET name = ?,
				description = ?,
				location = ?,
				datetime = ?
				WHERE ID = ?
				`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events
			WHERE ID = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}
	return nil
}
