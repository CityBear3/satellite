package driver

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func CreateDB(conf DBConfig) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := sql.Open(conf.Driver, dataSource)
	if err != nil {
		return nil, err
	}
	boil.SetDB(db)

	return db, nil
}