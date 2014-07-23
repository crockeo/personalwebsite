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

type DisplayProject struct {
	Title           string   // The title of the project
	FirstScreenshot string   // The first screenshot
	Screenshots     []string // The rest of the screenshots
	Language        string   // The language the project was written in
	ShortDesc       string   // A short version of the description
	Description     string   // A description of the project
}

func toDisplayProject(project *database.Project) *DisplayProject {
	var firstscreenshot string
	var screenshots []string

	if project.Screenshots == nil {
		firstscreenshot = ""
		screenshots = nil
	} else {
		firstscreenshot = project.Screenshots[0]

		if len(project.Screenshots) > 1 {
			screenshots = project.Screenshots[1:]
		} else {
			screenshots = nil
		}
	}

	return &DisplayProject{
		Title:           project.Title,
		FirstScreenshot: firstscreenshot,
		Screenshots:     screenshots,
		Language:        project.Language,
		ShortDesc:       project.ShortDesc,
		Description:     project.Description,
	}
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
				ErrorHandler(w, r, 404)
			} else {
				helpers.SendPage(w, "course", course)
			}
		}
	}
}

// Displaying an individual project
func ProjectProjectHandler(w http.ResponseWriter, r *http.Request) {
	root := "/project/project/"
	title := r.URL.Path[len(root):]

	if title == "" {
		ErrorHandler(w, r, 404)
	} else {
		db, err := database.OpenDefaultDatabase()

		if err != nil {
			ErrorHandler(w, r, 503)
		} else {
			project, err := database.GetProjectByTitle(db, title)

			if err != nil {
				ErrorHandler(w, r, 404)
			} else {
				helpers.SendPage(w, "project", toDisplayProject(project))
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
