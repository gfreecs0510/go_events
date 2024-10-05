package models

import (
	"gfreecs0510/events/src/clients"
)

type Event struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	CreatedAt   string `json:"created_at"`
	UserId      int64  `json:"user_id"`
}

func (e *Event) Create() error {
	q := `
		INSERT INTO events (name, description, location, created_at, user_id)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP, ?)
	`

	stmt, err := clients.DB.Prepare(q)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.UserId)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id
	return nil
}

func (e *Event) Update() error {
	q := `UPDATE events SET name = ?, description = ?, location = ?, user_id = ? WHERE id = ?`

	stmt, err := clients.DB.Prepare(q)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.UserId, e.ID)
	return err
}

func (e *Event) Delete() error {
	q := `DELETE FROM events WHERE id = ?`

	stmt, err := clients.DB.Prepare(q)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func GetAllEvents() ([]Event, error) {
	var events = []Event{}
	rows, err := clients.DB.Query("SELECT * FROM events")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event Event
		rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.CreatedAt, &event.UserId)
		events = append(events, event)
	}

	return events, nil
}

func GetEventViaId(id int64) (Event, error) {
	var event = Event{}

	q := `SELECT * FROM events WHERE id = ?`

	row := clients.DB.QueryRow(q, id)

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.CreatedAt, &event.UserId)

	return event, err
}
