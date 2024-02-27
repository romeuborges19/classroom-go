package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Group struct {
	ID          int
	Name        string
	Classes     []*Class
	Students    []*Student 
	CreatedAt   *time.Time
	ModifiedAt  *time.Time
}

type Class struct {
	ID    string
	Name  string
}

type Student struct {
	ID 		 int     `json:"id"`
	FullName string  `json:"fullname"`
	Email 	 string  `json:"email"`
}

func (s *Student) Value() (driver.Value, error) {
	return json.Marshal(s)
}
