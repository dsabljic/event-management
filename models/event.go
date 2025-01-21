package models

import (
	"time"

	"github.com/dsabljic/event-management/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

type Registration struct {
	ID        int64
	EventName string
	UserEmail string
}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, dateTime, user_id)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&e.ID)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events`
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

func GetEventById(id int64) (*Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events WHERE id = $1`
	var event Event
	err := db.DB.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = $1, description = $2, location = $3, dateTime = $4
		WHERE id = $5
	`
	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id = $1`
	_, err := db.DB.Exec(query, event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := `INSERT INTO registrations (event_id, user_id) VALUES ($1, $2)`
	_, err := db.DB.Exec(query, e.ID, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = $1 AND user_id = $2`
	_, err := db.DB.Exec(query, e.ID, userId)
	return err
}

func FetchRegistrations(eventId int64) ([]Registration, error) {
	query := `
		SELECT r.id, e.name, u.email 
		FROM registrations r
		JOIN events e ON r.event_id = e.id
		JOIN users u ON u.id = r.user_id
		WHERE r.event_id = $1
	`
	rows, err := db.DB.Query(query, eventId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registrations []Registration
	for rows.Next() {
		var registration Registration
		err := rows.Scan(&registration.ID, &registration.EventName, &registration.UserEmail)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}

	return registrations, nil
}
