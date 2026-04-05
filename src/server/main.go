package main

import (
	"log"
	"net/http"
	"os"

	"easy-login/server/application"
	"easy-login/server/infrastructure/sqlite"
	"easy-login/server/infrastructure/system"
	httpapi "easy-login/server/interfaces/http"
)

func main() {
	databasePath := os.Getenv("SQLITE_PATH")
	db, err := sqlite.OpenDatabase(databasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store, err := sqlite.NewStore(db)
	if err != nil {
		log.Fatal(err)
	}

	generator := system.RandomHexGenerator{}

	service := application.NewService(
		store,
		store,
		generator,
		generator,
	)

	handler := httpapi.NewHandler(service)

	address := ":8080"
	if value := os.Getenv("PORT"); value != "" {
		address = ":" + value
	}

	log.Printf("easy-login server listening on %s", address)
	log.Fatal(http.ListenAndServe(address, handler.Routes()))
}
