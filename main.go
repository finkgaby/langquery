package main

import (
	"github.com/joho/godotenv"
	"langquery/commons"
	"langquery/server"
	"log"
	"os"
)

func main() {
	loadConfigurations()
	server.ConnectNatsAndSubscribe()
}

func loadConfigurations() {
	log.Printf("Loading configuration")

	env := os.Getenv("env")
	if env == "local" {
		err := godotenv.Load(".env.local", ".env")
		commons.CheckErr(err)

	} else {
		err := godotenv.Load()
		commons.CheckErr(err)
	}
}
