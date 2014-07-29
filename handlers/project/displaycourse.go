package project

import (
	"github.com/crockeo/personalwebsite/database"
	"html/template"
	"strings"
)

// The displayCourse type
type displayCourse struct {
	SerTitle    string
	Title       string
	Inst        string
	Description template.HTML
	Comments    template.HTML
}

// Converting a Course to a displayCourse
func toDisplayCourse(course *database.Course) *displayCourse {
	fixedDescription := template.HTML(strings.Replace(course.Description, "\n", "<br>", -1))
	fixedComments := template.HTML(strings.Replace(course.Comments, "\n", "<br>", -1))

	return &displayCourse{
		SerTitle:    course.SerTitle,
		Title:       course.Title,
		Inst:        course.Inst,
		Description: fixedDescription,
		Comments:    fixedComments,
	}
}
