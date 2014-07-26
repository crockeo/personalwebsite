package project

import "github.com/crockeo/personalwebsite/database"

// The displayProject type
type displayProject struct {
	Title           string   // The title of the project
	FirstScreenshot string   // The first screenshot
	Screenshots     []string // The rest of the screenshots
	Language        string   // The language the project was written in
	ShortDesc       string   // A short version of the description
	Description     string   // A description of the project
}

// Converting a Project to a displayProject
func toDisplayProject(project *database.Project) *displayProject {
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

	return &displayProject{
		Title:           project.Title,
		FirstScreenshot: firstscreenshot,
		Screenshots:     screenshots,
		Language:        project.Language,
		ShortDesc:       project.ShortDesc,
		Description:     project.Description,
	}
}
