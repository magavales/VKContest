package main

import (
	"VK"
	"VK/pkg/handler"
	"log"
)

func main() {
	router := new(handler.Handler).InitRouter()

	server := new(AuthPage.Server)
	err := server.InitServer("8080", router)
	if err != nil {
		log.Fatalf("Server can't be opened: %s", err)
	}
}
