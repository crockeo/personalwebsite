package database

import (
	"database/sql"
	"strings"
)

func makeProjectTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS projects (id INTEGER NOT NULL PRIMARY KEY, title TEXT NOT NULL, screenshots TEXT NOT NULL, language TEXT NOT NULL, desc TEXT NOT NULL)")

	return err
}

type Project struct {
	Id          int      // The id of the project
	Title       string   // The title of the project
	Screenshots []string // URLs to screenshots of the project in action
	Language    string   // The language the project was written in
	ShortDesc   string   // A short version of the description
	Desc        string   // A description of the project
}

// Making a new project object
func MakeProject(id int, title string, screenshots []string, language string, shortdesc string, desc string) *Project {
	return &Project{
		Id:          id,
		Title:       title,
		Screenshots: screenshots,
		Language:    language,
		ShortDesc:   shortdesc,
		Desc:        desc,
	}
}

// Making a list of projects from a sql.Rows object
func makeProjects(rows *sql.Rows) ([]*Project, error) {
	projects := make([]*Project, 0)

	var id int
	var title string
	var screenshots string
	var language string
	var shortdesc string
	var desc string

	for rows.Next() {
		err := rows.Scan(&id, &title, &screenshots, &language, &shortdesc, &desc)

		if err != nil {
			return nil, err
		}

		projects = append(projects, MakeProject(id, title, strings.Split(screenshots, ","), language, shortdesc, desc))
	}

	return projects, nil
}

// Getting all of the projects
func GetProjects(db *sql.DB) ([]*Project, error) {
	rows, err := db.Query("SELECT * FROM projects")

	if err != nil {
		return nil, err
	}

	return makeProjects(rows)
}

// Querying rows with a 'WHERE' statement
func queryProjectWithWhere(db *sql.DB, field string, value string) ([]*Project, error) {
	stmt, err := db.Prepare("SELECT * FROM projects WHERE $1 = $2")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(field, value)

	if err != nil {
		return nil, err
	}

	return makeProjects(rows)
}

// Querying a project by language
func QueryProjectByLanguage(db *sql.DB, language string) ([]*Project, error) {
	return queryProjectWithWhere(db, "language", language)
}
