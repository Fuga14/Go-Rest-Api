package models

import (
	"fmt"
	"rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `bind:"required"`
	Description string    `bind:"required"`
	Location    string    `bind:"required"`
	DateTime    time.Time `bind:"required"`
	UserID      int64
}

var events []Event

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	fmt.Println(event)

	return &event, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)

	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(e.ID)

	return err
}
