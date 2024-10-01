package models

import (
	"log"

	"github.com/gobuffalo/pop/v6"
)

var DB *pop.Connection

func init() {
	var err error
	DB, err = pop.Connect("production")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}
