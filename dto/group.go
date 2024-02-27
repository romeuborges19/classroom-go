package dto

import "time"

type Group struct {
	ID        int
	Name      string
	Classes   []ClassInfo
	Students  []StudentInfo
	CreatedAt *time.Time
}

type CreateGroup struct {
	Name    string
	Classes string
}

type ClassInfo struct {
	ID         string    `json:"id"`
	CourseName string `json:"courseName"`
}

type StudentInfo struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type CourseInfo struct {
	CourseName  string          `json:"courseName"`
	Students    *[]StudentInfo  `json:"students"`
}
