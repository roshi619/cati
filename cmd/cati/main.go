package main

import (
	"log"

	"github.com/roshi619/cati/internal/command"
)

func main() {
	if err := command.Root.Execute(); err != nil {
		log.Fatal(err)
	}
}
