package main

import (
	"log"

	"github.com/divakarpatil51/go_social/internal/db"
	"github.com/divakarpatil51/go_social/internal/env"
	"github.com/divakarpatil51/go_social/internal/store"
)

func main() {
	config := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:               env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConnections: env.GetInt("MAX_OPEN_CONNECTIONS", 30),
			maxIdleConnections: env.GetInt("MAX_IDLE_CONNECTIONS", 30),
		},
	}

	db, err := db.New(config.db.addr, config.db.maxOpenConnections, config.db.maxIdleConnections)
	if err != nil {
		log.Panic(err)
	}

	store := store.NewStorage(db)
	app := &application{
		config: config,
		store:  store,
	}

	mount := app.mount()
	log.Fatal(app.run(mount))
}
