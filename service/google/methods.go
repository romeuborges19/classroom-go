package service

import (
	"classroom/dto"
	"log"
)



func (g *googleService) GetCourses(respch chan map[string]string){
	r, err := g.classroom.Courses.List().PageSize(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve courses. %v", err)
	}

	resp := make(map[string]string)

	for _, c := range r.Courses {
		resp[c.Id] = c.Name
	}

	respch <- resp
}

func (g *googleService) GetCourseData(courseId string, ch chan []dto.StudentInfo){
	_, err := g.classroom.Courses.Get(courseId).Do()
	if err != nil {
		log.Fatal(err)
	}

	var studentsData []dto.StudentInfo
	nextPageToken := ""
	index := 0
	for {
		call := g.classroom.Courses.Students.List(courseId).PageToken(nextPageToken)
		res, err := call.Do()
		if err != nil {
			log.Fatal(err)
		}

		for _, student := range res.Students {
			studentsData = append(studentsData, dto.StudentInfo{
				ID: index,
				FullName: student.Profile.Name.FullName,
				Email: student.Profile.EmailAddress,
			})
			index = index + 1
		}

		nextPageToken = res.NextPageToken
		if nextPageToken == "" {
			break
		}
	}
	ch <- studentsData
}
