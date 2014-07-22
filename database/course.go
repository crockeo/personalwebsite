package database

import (
	"database/sql"
	"strconv"
)

func makeCourseTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS courses (id INTEGER NOT NULL PRIMARY KEY, title TEXT NOT NULL, sertitle TEXT UNIQUE NOT NULL, inst TEXT NOT NULL, desc TEXT NOT NULL)")

	return err
}

type Course struct {
	Id       int    // The (local) id of the course
	Title    string // The title of the course
	SerTitle string // The serialized version of the course title
	Inst     string // The institution that offers the course
	Desc     string // The description of the course
}

// Making a new course object
func MakeNewCourse(id int, title string, sertitle string, inst string, desc string) *Course {
	return &Course{
		Id:       id,
		Title:    title,
		SerTitle: sertitle,
		Inst:     inst,
		Desc:     desc,
	}
}

// Getting all of the courses
func GetCourses(db *sql.DB) ([]*Course, error) {
	rows, err := db.Query("SELECT * FROM courses ORDER BY id DESC")

	if err != nil {
		return nil, err
	}

	courses := make([]*Course, 0)

	var id int
	var title string
	var sertitle string
	var inst string
	var desc string

	for rows.Next() {
		err = rows.Scan(&id, &title, &sertitle, &inst, &desc)

		if err != nil {
			return nil, err
		}

		courses = append(courses, MakeNewCourse(id, title, sertitle, inst, desc))
	}

	return courses, nil
}

// Querying a row with a where
func queryWithWhere(db *sql.DB, field string, value string) (*Course, error) {
	stmt, err := db.Prepare("SELECT * FROM courses WHERE $1 = $2")

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(field, value)

	if row == nil {
		return nil, RowDoesNotExistError
	}

	var id int
	var title string
	var sertitle string
	var inst string
	var desc string

	err = row.Scan(&id, &title, &sertitle, &inst, &desc)

	if err != nil {
		return nil, err
	}

	return MakeNewCourse(id, title, sertitle, inst, desc), nil
}

// Getting a course by its ID
func GetCourseByID(db *sql.DB, id int) (*Course, error) {
	return queryWithWhere(db, "id", strconv.FormatInt(int64(id), 10))
}

// Getting a course by its serialized title
func GetCourseBySerTitle(db *sql.DB, sertitle string) (*Course, error) {
	return queryWithWhere(db, "sertitle", sertitle)
}
