package models

import (
	"database/sql"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash []byte
}

func GetUserByUsername(db *sql.DB, username string) (User, error) {
	var user = User{}

	rows, err := db.Query(
		"select id, username, email, password from users where username=?",
		username,
	)

	if err != nil {
		return user, errors.Wrap(err, "database query error: ")
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash); err != nil {
			return user, errors.Wrap(err, "database scan error: ")
		}
	}
	if err := rows.Err(); err != nil {
		return user, errors.Wrap(err, "database rows error: ")
	}

	return user, nil
}

func CreateUser(db *sql.DB, username string, password string, email string) (User, error) {
	var user = User{}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
        return user, errors.Wrap(err, "hash password error:")
    }
	encodePW := string(hash)  // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可


	stmt, err := db.Prepare("INSERT users SET username=?, email=?, password=?")
	if err != nil {
		return user, errors.Wrap(err, "prepare db error:")
	}

	res, err := stmt.Exec(username, email, encodePW)
	if err != nil {
		return user, errors.Wrap(err, "insert db error:")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return user, errors.Wrap(err, "last insert id error:")
	}

	user.ID = int(id)
	user.Username = username
	user.Email = email
	user.PasswordHash = []byte(encodePW)

	return user, nil
}
