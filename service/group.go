package service

import (
	"classroom/dto"
	"classroom/model"
	"classroom/repository"
	service "classroom/service/google"
	"database/sql"
	"strings"
)

type GroupService interface {
	CreateGroup (group dto.CreateGroup, db *sql.DB) (int64, error)
}

type groupService struct {
	googleService service.GoogleService
	dao repository.DAO
}

func NewGroupService(googleService service.GoogleService, dao repository.DAO) GroupService {
	return &groupService{
		googleService: googleService,
		dao: dao,
	}
}

func (g *groupService) CreateGroup (group dto.CreateGroup, db *sql.DB) (int64, error) {
	coursesID := strings.Split(group.Classes, ",")
	
	respch := make(chan []dto.CourseInfo)
	go g.googleService.GetListOfCoursesData(coursesID, respch)
	groupInfo := <- respch

	var courses []model.Course
	for _, course := range groupInfo {
		var students []*model.Student
		for _, student := range *course.Students {
			studentModel := &model.Student{
				ID: student.ID,
				FullName: student.FullName,
				Email: student.Email,
			}

			students = append(students, studentModel)
		}

		courseModel := model.Course{
			ID: course.ID,
			Name: course.CourseName,
			Students: students,
		}
		courses = append(courses, courseModel)	
	}

	groupModel := model.Group{
		Name: group.Name,
		Classes: courses,
	}

	groupID, err := g.dao.NewGroupQuery().CreateGroup(&groupModel, db)
	if err != nil {
		return 0, err 
	}

	return groupID, nil
}
