package database

import (
	"github.com/jmoiron/sqlx"
)

func makeCourseTable(db *sqlx.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS courses (sertitle TEXT NOT NULL PRIMARY KEY, title TEXT NOT NULL, inst TEXT NOT NULL, description TEXT NOT NULL, comments TEXT NOT NULL)")

	return err
}

type Course struct {
	Title       string // The title of the course
	SerTitle    string // The serialized version of the course title
	Inst        string // The institution that offers the course
	Description string // The description of the course
	Comments    string // My personal comments on the course
}

// Making a new course object
func MakeNewCourse(sertitle string, title string, inst string, description string, comments string) *Course {
	return &Course{
		SerTitle:    sertitle,
		Title:       title,
		Inst:        inst,
		Description: description,
		Comments:    comments,
	}
}

// Getting all of the courses
func GetCourses(db *sqlx.DB) ([]*Course, error) {
	rows, err := db.Query("SELECT * FROM courses")

	if err != nil {
		return nil, err
	}

	courses := make([]*Course, 0)

	var title string
	var sertitle string
	var inst string
	var description string
	var comments string

	for rows.Next() {
		err = rows.Scan(&sertitle, &title, &inst, &description, &comments)

		if err != nil {
			return nil, err
		}

		courses = append(courses, MakeNewCourse(sertitle, title, inst, description, comments))
	}

	return courses, nil
}

// Quickly getting all of the courses
func QuickGetCourses() []*Course {
	db := QuickOpenDB()
	defer db.Close()

	courses, err := GetCourses(db)

	if err != nil {
		panic(err)
	}

	return courses
}

// Getting a course by its serialized title
func GetCourseBySerTitle(db *sqlx.DB, sertitle string) (*Course, error) {
	stmt, err := db.Prepare("SELECT * FROM courses WHERE sertitle = $1")

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(sertitle)

	if row == nil {
		return nil, RowDoesNotExistError
	}

	var title string
	var nsertitle string
	var inst string
	var description string
	var comments string

	err = row.Scan(&nsertitle, &title, &inst, &description, &comments)

	if err != nil {
		return nil, err
	}

	return MakeNewCourse(nsertitle, title, inst, description, comments), nil
}

// Quickly getting a course by its serialized title
func QuickGetCourseBySerTitle(sertitle string) *Course {
	db := QuickOpenDB()
	defer db.Close()

	course, err := GetCourseBySerTitle(db, sertitle)

	if err != nil {
		panic(err)
	}

	return course
}

// Inserting a course
func InsertCourse(db *sqlx.DB, course *Course) error {
	stmt, err := db.Prepare("INSERT INTO courses(sertitle, title, inst, description, comments) values($1, $2, $3, $4, $5)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(course.SerTitle, course.Title, course.Inst, course.Description, course.Comments)

	return err
}

// Quickly inserting a course
func QuickInsertCourse(course *Course) {
	db := QuickOpenDB()
	defer db.Close()

	err := InsertCourse(db, course)

	if err != nil {
		panic(err)
	}
}
