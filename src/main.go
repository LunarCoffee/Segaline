package main

import (
	"log"
	"os"
	"segaline/src/server"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("No file root directory or template root directory was provided!")
		return
	}

	fileServer := server.NewFileServer(os.Args[1], os.Args[2])
	if err := fileServer.Start("0.0.0.0:1440"); err != nil {
		log.Fatalln("An error occurred while starting the server!")
	}
}
