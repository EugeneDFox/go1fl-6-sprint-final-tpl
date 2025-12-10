package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags|log.Lshortfile)
	s := server.NewServer(logger)

	log.Fatal(s.Start())
}
