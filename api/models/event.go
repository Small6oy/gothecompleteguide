package models

import (
	"fmt"
	"time"

	"example.com/api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, userID)
		VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Failed to Prepare Statement", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		fmt.Println("Failed to Exec Statement", err)
		return err
	}

	fmt.Println(result)
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

func GetEventById(eventId int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, eventId)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e Event) Remove() {
	events = append(events, e)
}
