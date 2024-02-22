package main

import (
	"log"

	"classroom/api"
)

func main() {
	s := api.NewServer()
	log.Fatal(s.Start())
}
