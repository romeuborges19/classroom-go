package dto

import "time"

type Group struct {
	ID        int
	Name      string
	Classes   []CourseInfo
	CreatedAt *time.Time
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
