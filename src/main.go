package main

import (
	"log"
	"segaline/src/server"
)

func main() {
	fileServer := server.NewFileServer("resources/www")
	if err := fileServer.Start("0.0.0.0:1440"); err != nil {
		log.Fatalln("An error occurred while starting the server!")
	}
}
