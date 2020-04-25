package main

import (
	"fmt"
	"log"
	"os"
	"segaline/src/server"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: " + os.Args[0] + " <static file root> <template root>")
	} else {
		fileServer := server.NewFileServer(os.Args[1], os.Args[2])
		if err := fileServer.Start("0.0.0.0:1440"); err != nil {
			log.Fatalln("An error occurred while starting the server!")
		}
	}
}
