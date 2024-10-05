package models

import "gfreecs0510/events/src/clients"

type Registration struct {
	ID        int64
	UserId    int64  `json:"user_id"`
	EventId   int64  `json:"event_id"`
	CreatedAt string `json:"created_at"`
}

func (r *Registration) Create() error {
	q := `
		INSERT INTO registrations (event_id, user_id, created_at)
		VALUES(?, ?, CURRENT_TIMESTAMP)
	`

	stmt, err := clients.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(r.EventId, r.UserId)

	if err != nil {
		return err
	}

	r.ID, err = res.LastInsertId()

	return err
}

func (r *Registration) Delete() error {
	q := `
		DELETE FROM registrations WHERE event_id = ? AND user_id = ?;
	`

	stmt, err := clients.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(r.EventId, r.UserId)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()

	return err
}

func GetRegistration(userID, eventId int64) (Registration, error) {
	var reg Registration

	q := `SELECT id, user_id, event_id, created_at FROM registrations WHERE user_id = ? AND event_id = ?`

	row := clients.DB.QueryRow(q, userID, eventId)

	err := row.Scan(&reg.ID, &reg.UserId, &reg.EventId)

	return reg, err
}
