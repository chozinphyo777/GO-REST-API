package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{} // slice of Event structs
func (e Event) Save() {
	//later save to database
	events = append(events, e)

}
