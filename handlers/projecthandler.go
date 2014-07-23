package handlers

import (
	"fmt"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/helpers"
	"net/http"
)

type projectRootData struct {
	Courses  []*database.Course  // The list of courses
	Projects []*database.Project // The list of projects
}

// Displaying an individual course
func ProjectCourseHandler(w http.ResponseWriter, r *http.Request) {
	root := "/project/course/"
	sertitle := r.URL.Path[len(root):]

	if sertitle == "" {
		ErrorHandler(w, r, 404)
	} else {
		db, err := database.OpenDefaultDatabase()

		if err != nil {
			ErrorHandler(w, r, 503)
		} else {
			course, err := database.GetCourseBySerTitle(db, sertitle)

			if err != nil {
				fmt.Println(err.Error())
				ErrorHandler(w, r, 404)
			} else {
				helpers.SendPage(w, "course", course)
			}
		}
	}
}

// The base ProjectHandler
func loadProjectRootData() (*projectRootData, error) {
	db, err := database.OpenDefaultDatabase()

	if err != nil {
		return nil, err
	}

	courses, err := database.GetCourses(db)

	if err != nil {
		return nil, err
	}

	projects, err := database.GetProjects(db)

	if err != nil {
		return nil, err
	}

	db.Close()

	return &projectRootData{
		Courses:  courses,
		Projects: projects,
	}, nil
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/project/" {
		ErrorHandler(w, r, 404)
	} else {
		data, err := loadProjectRootData()

		if err != nil {
			fmt.Println(err.Error())
			ErrorHandler(w, r, 503)
		} else {
			helpers.SendPage(w, "projectroot", data)
		}
	}
}
