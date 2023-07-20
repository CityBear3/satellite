package main

import (
	"database/sql"
	"github.com/CityBear3/satellite/driver"
	"github.com/CityBear3/satellite/pkg/env"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	dbHost := env.GetStrEnv("MYSQL_HOST", "localhost")
	dbPort, err := env.GetIntEnv("MYSQL_PORT", 3306)
	if err != nil {
		log.Fatalln(err)
	}
	dbName := env.GetStrEnv("MYSQL_DBNAME", "satellite")
	dbUser, err := env.GetRequiredStrEnv("MYSQL_USER")
	if err != nil {
		log.Fatalln(err)
	}
	dbPass, err := env.GetRequiredStrEnv("MYSQL_PASS")
	if err != nil {
		log.Fatalln(err)
	}
	dbConfig := driver.DBConfig{
		Driver:   "mysql",
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPass,
		DbName:   dbName,
	}

	db, err := driver.CreateDB(dbConfig)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	port, err := env.GetIntEnv("SERVER_PORT", 8080)
	if err != nil {
		log.Fatalln(err)
	}
	isDeveloping, err := env.GetBoolEnv("IS_DEVELOPING", false)
	if err != nil {
		log.Fatalln(err)
	}
	server := driver.NewServer(port, isDeveloping)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
