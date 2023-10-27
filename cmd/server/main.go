package main

import (
	"log"

	"github.com/CityBear3/satellite/internal/driver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := driver.Config{}
	if err := driver.LoadConfig(&cfg); err != nil {
		log.Fatalln(err)
	}

	server := driver.NewServer(cfg)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
