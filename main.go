package main

import (
	"log"

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

	log.Println(db)
	dao := repository.NewDAO()
	c := google.NewClassroomService()
	g := google.NewGoogleService(c)
	gr := service.NewGroupService(g, dao)

	s := routes.NewServer(g, gr, db)

	if err = db.Ping(); err != nil {
		log.Fatalf("treste %v", err)
	}


	log.Fatal(s.Start())
}
