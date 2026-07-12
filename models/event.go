package models

import (
	"time"

	"example.com/go-rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{} // slice of Event structs
func (e Event) Save() error {
	query := `INSERT INTO events (name, description, location, dateTime, user_id) Values (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query) // complex query, so we prepare it first to avoid SQL injection and improve performance
	if err != nil {
		return err
	}
	defer stmt.Close()                                                                // close the statement after execution
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) // execute the statement with the event data
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	e.ID = id

	// events = append(events, e)

	return err

}
func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query) // execute the query to get all events (Prepare is not needed for simple queries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event                                                                                               // declare a variable to hold the event data
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) // scan the row data into the event variable
		if err != nil {
			return nil, err
		}
		events = append(events, event) // append the event to the events slice
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id) // execute the query to get the event by ID (Prepare is not needed for simple queries)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) // scan the row data into the event variable
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events 
		SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ? 
		WHERE id = ?
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)
	if err != nil {
		return err
	}

	return nil

}
