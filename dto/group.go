package dto

import "time"

type Group struct {
	ID        int          `redis:"id"`
	Name      string       `redis:"name"`
	Classes   []CourseInfo `redis:"classes"`
	CreatedAt *time.Time   `redis:"createdAt"`
}

type CreateGroup struct {
	Name    string
	Classes string
}

type StudentInfo struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type CourseInfo struct {
	ID          string 			`json:"id"`
	CourseName  string          `json:"courseName"`
	Students    *[]StudentInfo  `json:"students,omitempty"`
}
