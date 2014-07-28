package schema

import "time"

// A post to be displayed on the blog
type Post struct {
	Id      int64     `db:"id"`      // The indexed ID of the blog post
	Title   string    `db:"title"`   // The title of the blog post
	Author  string    `db:"author"`  // The blog post's author
	Body    string    `db:"body"`    // The body of the blog post
	Written time.Time `db:"written"` // The time the post was written (in UnixNano)
}
