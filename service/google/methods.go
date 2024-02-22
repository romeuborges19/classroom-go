package service

import (
	"fmt"
	"log"
)

func (g *googleService) GetCourses() {
	r, err := g.classroom.Courses.List().PageSize(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve courses. %v", err)
	}
	if len(r.Courses) > 0 {
		fmt.Print("Courses:\n")
		for _, c := range r.Courses {
			fmt.Printf("%s (%s)\n", c.Name, c.Id)
		}
	} else {
		fmt.Print("No courses found.")
	}
}
