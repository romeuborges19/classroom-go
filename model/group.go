package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Group struct {
	ID          int
	Name        string
	Classes     Courses
	CreatedAt   *time.Time
	ModifiedAt  *time.Time
}

type Course struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Students []*Student `json:"students"`
}

type Student struct {
	ID 		 int     `json:"id"`
	FullName string  `json:"fullname"`
	Email 	 string  `json:"email"`
}

type Courses []Course

func (c *Courses) Value() (driver.Value, error) {
	return json.Marshal(c)
}
