package helper

import "database/sql"

func GetTestDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "satellite:satellite@tcp(localhost:3307)/satellite")
	if err != nil {
		return nil, err
	}

	return db, nil
}
