package dba

import (
	"github.com/jmoiron/sqlx"
)

func Get_password(db *sqlx.DB, username string) (string, error) {
	var password string

	err := db.Get(&password,
		`select passwd from login
            where username = $1`,
		username)

	return password, err
}
