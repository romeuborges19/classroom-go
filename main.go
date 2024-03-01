package main

import (
	"log"
	"net/http"

	"classroom/api/routes"
	"classroom/repository"
	"classroom/service"
	google "classroom/service/google"
)

func main() {
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cache := repository.NewCache()

	dao := repository.NewDAO()
	c := google.NewClassroomService()
	g := google.NewGoogleService(c)
	gr := service.NewGroupService(g, dao, cache)

	s := routes.NewServer(g, gr, db)

	if err = db.Ping(); err != nil {
		log.Fatalf("treste %v", err)
	}

	server := s.Start()
	http.ListenAndServe(":8080", server)
}
