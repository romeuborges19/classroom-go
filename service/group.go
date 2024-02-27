package service

import (
	"classroom/dto"
	"classroom/model"
	"classroom/repository"
	"database/sql"
)

type GroupService interface {

}

type groupService struct {
	dao repository.DAO
}

func (g *groupService) CreateGroup (group dto.Group, db *sql.DB) (int64, error) {
	var students []*model.Student
	for _, student := range group.Students {
		studentModel := &model.Student{
			ID: student.ID,
			FullName: student.FullName,
			Email: student.Email,
		}
		students = append(students, studentModel)
	}

	groupModel := model.Group{
		Name: group.Name,
		Students: students,
	}

	groupID, err := g.dao.NewGroupQuery().CreateGroup(groupModel, db)
	if err != nil {
		return 0, err
	}

	return groupID, err
}
