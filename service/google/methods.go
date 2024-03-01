package service

import (
	"classroom/dto"
	"log"
	"time"
)



func (g *googleService) GetCourses(respch chan map[string]string){
	r, err := g.classroom.Courses.List().Fields("courses/id","courses/name").PageSize(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve courses. %v", err)
	}

	resp := make(map[string]string)

	for _, c := range r.Courses {
		resp[c.Id] = c.Name
	}

	respch <- resp
}

func (g *googleService) GetCourseData(courseId string, ch chan *dto.CourseInfo){
	res, err := g.classroom.Courses.Get(courseId).Fields("name").Do()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	var studentsData []dto.StudentInfo
	
	c := &dto.CourseInfo{
		CourseName: res.Name,
		Students: &studentsData,
	}

	nextPageToken := ""
	index := 0
	for {
		call := g.classroom.Courses.Students.List(courseId).Fields(
			"students/profile/name/fullName", 
			"students/profile/emailAddress", 
			"nextPageToken",
			).PageToken(nextPageToken)
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
	ch <- c
}

func (g *googleService) GetListOfCoursesData(coursesId []string, ch chan []dto.CourseInfo){
	var coursesInfo []dto.CourseInfo
	for _, courseId := range coursesId {
		start := time.Now()
		res, err := g.classroom.Courses.Get(courseId).Fields("id","name").Do()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("API - Cursos: %v", time.Since(start))

		var studentsData []dto.StudentInfo
		
		c := dto.CourseInfo{
			ID: res.Id,
			CourseName: res.Name,
			Students: &studentsData,
		}

		nextPageToken := ""
		i := 0
		for {
			start = time.Now()
			call := g.classroom.Courses.Students.List(courseId).Fields(
			"students/profile/name/fullName", 
			"students/profile/emailAddress", 
			"nextPageToken",
			).PageToken(nextPageToken)
			res, err := call.Do()
			log.Printf("API - Estudantes: %v", time.Since(start))

			if err != nil {
				log.Fatal(err)
			}

			for _, student := range res.Students {
				studentsData = append(studentsData, dto.StudentInfo{
					ID: i,
					FullName: student.Profile.Name.FullName,
					Email: student.Profile.EmailAddress,
				})
				i++
			}

			nextPageToken = res.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
		coursesInfo = append(coursesInfo, c)	
	}
	ch <- coursesInfo
}
