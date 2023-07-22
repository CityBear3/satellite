package main

import (
	"database/sql"
	"github.com/CityBear3/satellite/driver"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	cfg, err := driver.LoadConfig("application.yml")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := driver.CreateDB(cfg.DBConfig)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	server := driver.NewServer(cfg.ServerConfig)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
