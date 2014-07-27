package schema

// Data to store for displaying and adding new courses
type Course struct {
	SerTitle    string `db:"ser_title"`   // The course's serialized title
	Title       string `db:"title"`       // The course's unserialized title
	Inst        string `db:"inst"`        // The institution that teaches the course
	Description string `db:"description"` // The institutions official description of the course
	Comments    string `db:"comments"`    // My personal comments on the course
}
