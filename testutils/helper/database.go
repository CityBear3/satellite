package helper

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func GetTestDB() (*sqlx.DB, error) {
	db, err := sql.Open("mysql", "satellite:satellite@tcp(localhost:3307)/satellite?parseTime=true")
	if err != nil {
		return nil, err
	}

	return sqlx.NewDb(db, "mysql"), nil
}
