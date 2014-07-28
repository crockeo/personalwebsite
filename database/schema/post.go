package schema

import (
	"github.com/coopernurse/gorp"
	"log"
	"time"
)

func toString(t time.Time) string { return t.Format(time.UnixDate) }
func fromString(st string) time.Time {
	t, err := time.Parse(time.UnixDate, st)

	if err != nil {
		log.Fatal(err)
	}

	return t
}

// A post to be displayed on the blog
type Post struct {
	Id       int64     `db:"id"`      // The indexed ID of the blog post
	Title    string    `db:"title"`   // The title of the blog post
	Author   string    `db:"author"`  // The blog post's author
	Body     string    `db:"body"`    // The body of the blog post
	SWritten string    `db:"written"` // The time the post was written (in string form)
	Written  time.Time `db:"-"`       // The time the post was written (in time form)
}

// gorp hooks
func (this *Post) PostGet(s gorp.SqlExecutor) {
	this.Written = fromString(this.SWritten)
}

func (this *Post) PreInsert(s gorp.SqlExecutor) {
	this.SWritten = toString(this.Written)
}

func (this *Post) PreUpdate(s gorp.SqlExecutor) {
	this.SWritten = toString(this.Written)
}
