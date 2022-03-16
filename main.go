package main

import (
	"github.com/henrique-sulimann/golang-restapi/server"
)

func main() {
	// database.StartDB()
	// database.StartMongo()
	server := server.NewServer()

	server.Run()
}
