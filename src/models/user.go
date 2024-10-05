package models

import "gfreecs0510/events/src/clients"

type User struct {
	ID       int64  `json:"id"`
	UserName string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Create() error {
	q := `
		INSERT INTO users (username, password) VALUES(? , ?)
	`

	stmt, err := clients.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.UserName, u.Password)

	return err
}

func GetUserViaUsername(username string) (User, error) {
	var user User
	q := `SELECT * FROM users WHERE username = ?`

	smt, err := clients.DB.Prepare(q)

	if err != nil {
		return user, err
	}

	defer smt.Close()

	row := smt.QueryRow(username)

	row.Scan(&user.ID, &user.UserName, &user.Password)

	return user, err
}
