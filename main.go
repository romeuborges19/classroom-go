package main

import (
	"log"

	"classroom/api"
	service "classroom/service/google"
)

func main() {
	c := service.NewClassroomService()
	g := service.NewGoogleService(c)

	s := api.NewServer(g)
	log.Fatal(s.Start())
}
