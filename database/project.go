package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
)

func makeProjectTable(db *sqlx.DB) error {
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

// Making a list of projects from a sqlx.Rows object
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
func GetProjects(db *sqlx.DB) ([]*Project, error) {
	rows, err := db.Query("SELECT * FROM projects")

	if err != nil {
		return nil, err
	}

	return makeProjects(rows)
}

// Quickly getting all of the projects
func QuickGetProjects() []*Project {
	db := QuickOpenDB()
	defer db.Close()

	projects, err := GetProjects(db)

	if err != nil {
		panic(err)
	}

	return projects
}

// Getting a project by its title
func GetProjectByTitle(db *sqlx.DB, title string) (*Project, error) {
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

// Quickly getting a project by its title
func QuickGetProjectByTitle(title string) *Project {
	db := QuickOpenDB()
	defer db.Close()

	project, err := GetProjectByTitle(db, title)

	if err != nil {
		panic(err)
	}

	return project
}

// Inserting a post
func InsertProject(db *sqlx.DB, project *Project) error {
	var joined string
	if len(project.Screenshots) == 0 {
		joined = ""
	} else {
		joined = project.Screenshots[0]
		for i := 1; i < len(project.Screenshots); i++ {
			joined += "," + project.Screenshots[i]
		}
	}

	exec := "INSERT INTO projects(title, screenshots, language, shortdesc, description) values($1, $2, $3, $4, $5)"

	_, err := db.Exec(exec, project.Title, joined, project.Language, project.ShortDesc, project.Description)
	return err
}

// Quickly inserting a post
func QuickInsertProject(project *Project) {
	db := QuickOpenDB()
	defer db.Close()

	if err := InsertProject(db, project); err != nil {
		panic(err)
	}
}
