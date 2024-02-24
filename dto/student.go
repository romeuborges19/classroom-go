package dto

type StudentInfo struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type CourseInfo struct {
	CourseName  string          `json:"courseName"`
	Students    *[]StudentInfo  `json:"students"`
}
