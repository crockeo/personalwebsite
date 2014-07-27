package schema

// A single screenshot
type Screenshot struct {
	Title string `db:"title"` // The matching project's title
	URL   string `db:"url"`   // The url of the screenshot
}

// Data to store for displaying and adding new projects
type Project struct {
	Title       string `db:"title"`       // The title of the project
	Language    string `db:"language"`    // The language that the project was written in
	ShortDesc   string `db:"short_desc"`  // A short version of the description
	Description string `db:"description"` // The full description
}

// The combined project and screenshots
type CombinedProject struct {
	Project                  // Inheriting Project values
	Screenshots []Screenshot // The list of screenshots
}
