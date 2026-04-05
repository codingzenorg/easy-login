package main

import (
	"log"
	"net/http"
	"os"

	"easy-login/server/application"
	"easy-login/server/infrastructure/memory"
	"easy-login/server/infrastructure/system"
	httpapi "easy-login/server/interfaces/http"
)

func main() {
	playerRepository := memory.NewPlayerRepository()
	deviceRepository := memory.NewDeviceRegistrationRepository()
	generator := system.RandomHexGenerator{}

	service := application.NewService(
		playerRepository,
		deviceRepository,
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
