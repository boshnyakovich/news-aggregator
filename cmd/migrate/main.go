package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/pressly/goose"
	"github.com/boshnyakovich/news-aggregator/config"
)

const (
	migrationPath = "/app/init/db/migrations"
)

func main() {
	var conf config.Config

	if err := conf.Parse(); err != nil {
		log.Fatalf("error parsing config: %s", err.Error())
	}

	connURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Pass, conf.Database.Name)

	db, err := sql.Open("postgres", connURL)
	if err != nil {
		log.Fatalf("error connecting to db: %s", err)
	}

	err = goose.Up(db, migrationPath)
	if err != nil {
		log.Fatalf("failed executing migrations DB: %v\n", err)
	}

	version, err := goose.GetDBVersion(db)
	if err != nil {
		log.Println("error getting version of migrations", err)
	}

	if err := db.Close(); err != nil {
		log.Fatalf("failed to close DB: %v\n", err)
	}

	log.Printf("exec migrations succeeded. version is %d", version)
}