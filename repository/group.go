package repository

import (
	"classroom/model"
	"database/sql"
	"errors"
	"log"
)

type GroupQuery interface {
	CreateGroup (group *model.Group, db *sql.DB) (int64, error)
	GetGroups (db *sql.DB) ([]model.Group, error)
	GetGroupById (id int, db *sql.DB) (model.Group, error) 
}

type groupQuery struct {}

func (g *groupQuery) CreateGroup (group *model.Group, db *sql.DB) (int64, error) {
	query := `INSERT INTO "group"("name", "students") VALUES ($1, $2) RETURNING "id";`

	res, err := db.Exec(query, group.Name, group.Students)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id, err
}

func (g *groupQuery) GetGroups (db *sql.DB) ([]model.Group, error) {
	query := `SELECT * FROM "groups" ORDER BY "id" DESC;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var groups []model.Group
	var group model.Group

	for rows.Next() {
		err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.Students,
			&group.CreatedAt,
			&group.ModifiedAt,
		)

		if err == nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, err	
	}

	return groups, nil
}

func (g *groupQuery) GetGroupById (id int, db *sql.DB) (model.Group, error) {
	query := `SELECT * FROM "groups" WHERE "id"=$1`

	var group model.Group

	err := db.QueryRow(query, id).Scan(
		&group.ID,
		&group.Name,
		&group.Students,
		&group.CreatedAt,
		&group.ModifiedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Group{}, errors.New("Group not found.")
		}
		return model.Group{}, err
	}

	return group, nil
}
