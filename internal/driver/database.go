package driver

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func CreateDB(conf DBConfig) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)

	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
