package main

import (
	"log"

	"github.com/eternalq/project-server/internal/api/server"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
