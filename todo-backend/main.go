package main

import (
	"backend/server"
	"log"
)

func main() {
	httpServer := server.NewServer(9000)
	if err := httpServer.StartServer(); err != nil {
		log.Fatalln(err)
	}
}
