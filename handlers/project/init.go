package project

import "github.com/go-martini/martini"

// Initializing the Project handlers
func Init(m *martini.ClassicMartini) {
	m.Get("/project/course/:name", ProjectCourseHandler)
	m.Get("/project/project/:name", ProjectProjectHandler)
	m.Get("/project", ProjectHandler)
}
