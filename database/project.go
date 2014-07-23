package database

import (
	"database/sql"
	"strings"
)

func makeProjectTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS projects (title TEXT NOT NULL PRIMARY KEY, screenshots TEXT NOT NULL, language TEXT NOT NULL, shortdesc TEXT NOT NULL, description TEXT NOT NULL)")

	return err
}

type Project struct {
	Title       string   // The title of the project
	Screenshots []string // URLs to screenshots of the project in action
	Language    string   // The language the project was written in
	ShortDesc   string   // A short version of the description
	Description string   // A description of the project
}

// Making a new project object
func MakeProject(title string, screenshots []string, language string, shortdesc string, description string) *Project {
	return &Project{
		Title:       title,
		Screenshots: screenshots,
		Language:    language,
		ShortDesc:   shortdesc,
		Description: description,
	}
}

// Making a list of projects from a sql.Rows object
func makeProjects(rows *sql.Rows) ([]*Project, error) {
	projects := make([]*Project, 0)
	var title string
	var screenshots string
	var language string
	var shortdesc string
	var description string

	for rows.Next() {
		err := rows.Scan(&title, &screenshots, &language, &shortdesc, &description)

		if err != nil {
			return nil, err
		}

		projects = append(projects, MakeProject(title, strings.Split(screenshots, ","), language, shortdesc, description))
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

// Getting a project by its title
func GetProjectByTitle(db *sql.DB, title string) (*Project, error) {
	stmt, err := db.Prepare("SELECT * FROM projects WHERE title = $1")

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(title)

	var ntitle string
	var sscreenshots string
	var language string
	var shortdesc string
	var description string

	err = row.Scan(&ntitle, &sscreenshots, &language, &shortdesc, &description)

	var screenshots []string
	if sscreenshots == "" {
		screenshots = nil
	} else {
		screenshots = strings.Split(sscreenshots, ",")
	}

	return &Project{
		Title:       ntitle,
		Screenshots: screenshots,
		Language:    language,
		ShortDesc:   shortdesc,
		Description: description,
	}, nil
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
