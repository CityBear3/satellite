package driver

import (
	"database/sql"
	"fmt"
)

func CreateDB(conf DBConfig) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)

	db, err := sql.Open(conf.Driver, dataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
