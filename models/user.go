package models

import (
	"database/sql"

	"github.com/pkg/errors"
)

type User struct {
	ID           int
	Username     string
	PasswordHash []byte
}

func GetUserByUsername(db *sql.DB, username string) (User, error) {
	var user = User{}

	rows, err := db.Query(
		"select id, username, password from users where username=?",
		username,
	)

	if err != nil {
		return user, errors.Wrap(err, "database query error: ")
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.PasswordHash); err != nil {
			return user, errors.Wrap(err, "database scan error: ")
		}
	}
	if err := rows.Err(); err != nil {
		return user, errors.Wrap(err, "database rows error: ")
	}

	return user, nil
}
