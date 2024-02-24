package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Group struct {
	ID          int
	name        string
	students    []*Student 
	createdAt   *time.Time
	modifiedAt  *time.Time
}

type Student struct {
	ID 		 int     `json:"id"`
	FullName string  `json:"fullname"`
	Email 	 string  `json:"email"`
}

func (s *Student) Value() (driver.Value, error) {
	return json.Marshal(s)
}
