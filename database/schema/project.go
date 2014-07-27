package schema

import (
	"github.com/coopernurse/gorp"
	"strings"
)

func toList(sscreenshots string) []string { return strings.Split(sscreenshots, ",") }
func fromList(screenshots []string) string {
	if len(screenshots) == 0 {
		return ""
	}

	sscreenshots := screenshots[0]
	for i := 1; i < len(screenshots); i++ {
		sscreenshots += "," + screenshots[i]
	}

	return sscreenshots
}

// Data to store for displaying and adding new projects
type Project struct {
	Title        string   `db:"title"`       // The title of the project
	Language     string   `db:"language"`    // The language that the project was written in
	SScreenshots string   `db:"screenshots"` // The screenshots, in string form
	ShortDesc    string   `db:"short_desc"`  // A short version of the description
	Description  string   `db:"description"` // The full description
	Screenshots  []string `db:"-"`           // The list of strings that function as the screenshots
}

// gorp hooks
func (this *Project) PostGet(s gorp.SqlExecutor) error {
	this.Screenshots = toList(this.SScreenshots)
	return nil
}

func (this *Project) PreInsert(s gorp.SqlExecutor) error {
	this.SScreenshots = fromList(this.Screenshots)
	return nil
}

func (this *Project) PreUpdate(s gorp.SqlExecutor) error {
	this.SScreenshots = fromList(this.Screenshots)
	return nil
}
