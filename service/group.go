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
	cache repository.Cache
}

func NewGroupService(googleService service.GoogleService, dao repository.DAO, cache repository.Cache) GroupService {
	return &groupService{
		googleService: googleService,
		dao: dao,
		cache: cache,
	}
}

func (g *groupService) GetGroupById (id string, db *sql.DB) (dto.Group, error) {
	// Verifica se o grupo está salvo no Redis
	group, err := g.cache.GetGroup(id)
	if err == nil {
		groupDTO := dto.Group{group}
		return group, nil
	}

	return dto.Group{}, nil
}

func (g *groupService) CreateGroup (group dto.CreateGroup, db *sql.DB) (int64, error) {
	coursesID := strings.Split(group.Classes, ",")
	
	// Obtém dados sobre o grupo e armazena em model
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

	// Armazena dados do grupo no banco de dados
	groupID, err := g.dao.NewGroupQuery().CreateGroup(&groupModel, db)
	if err != nil {
		return 0, err 
	}

	groupModel.ID = int(groupID)

	err = g.cache.SetGroup(&groupModel)
	if err != nil {
		return 0, err
	}

	return groupID, nil
}
