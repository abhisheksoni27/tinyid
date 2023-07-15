package main

import (
	"log"

	"github.com/abhisheksoni27/tinyid"
)

func main() {
	id := tinyid.New(10)
	log.Printf("tinyid = %s", id)
}
