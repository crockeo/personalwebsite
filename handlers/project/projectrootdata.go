package project

import "github.com/crockeo/personalwebsite/database"

type projectRootData struct {
	Courses  []*database.Course  // The list of courses
	Projects []*database.Project // The list of projects
}

// The base ProjectHandler
func loadProjectRootData() (*projectRootData, error) {
	db := database.QuickOpenDB()

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
