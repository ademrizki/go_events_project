package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {

	query := `
		INSERT INTO events(name, description, location, dateTime, user_id) 
		VALUES (?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		panic("Prepare Save Event failed")
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		&e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID,
	)

	if err != nil {
		panic("Execute Save Event failed")
	}

	id, err := result.LastInsertId()

	e.ID = id

	return err

}

func GetEvents() ([]Event, error) {

	query := "SELECT * FROM events"

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

func GetEvent(id int64) (*Event, error) {

	query := "SELECT * FROM events WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
