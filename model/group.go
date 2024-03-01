package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Group struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Classes     Courses    `json:"classes"`
	CreatedAt   *time.Time `json:"createdAt"`
	ModifiedAt  *time.Time `json:"modifiedAt"`
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

func (g *Group) Value() (driver.Value, error) {
	return json.Marshal(g)
}
