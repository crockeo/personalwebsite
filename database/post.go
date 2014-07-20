package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/crockeo/personalwebsite/helpers"
	"html/template"
	"time"
)

// The post type
type Post struct {
	Id      int       // The ID of the post
	Title   string    // The title of the post
	Author  string    // The author of the post
	Body    string    // The body of the post
	Written time.Time // The time the post was written
}

// Converting the post to a string
func (post *Post) String() string {
	return fmt.Sprintf("id=%d;title=%s;author=%s;body=%s;written=%s;", post.Id, post.Title, post.Author, post.Body, post.Written.Format(time.UnixDate))
}

// Displaying the post for listing on the /blob/ page
func (post *Post) Display() template.HTML {
	return template.HTML(fmt.Sprintf(`
	<div class="col-md-6 col-md-offset-3">
		<div class="text-center"><a class="blog-link" href="/blog/%d">%s</a></div>

		<div class="text-center blog-author">%s</div>

		<div class="text-justify blog-content">
			%s
		</div>

		<div class="text-center blog-time">%s</div>
	</div>
	`, post.Id, post.Title, post.Author, helpers.ParseMarkdown(post.Body), post.Written.Format(time.UnixDate)))
}

// Making a new post
func MakeNewPostRaw(id int, title string, author string, body string, written time.Time) *Post {
	return &Post{
		Id:      id,
		Title:   title,
		Author:  author,
		Body:    body,
		Written: written,
	}
}

// Making a new post, save id
func MakeNewPost(title string, author string, body string) *Post {
	return MakeNewPostRaw(1, title, author, body, time.Now())
}

// Getting every post
func GetPosts(db *sql.DB) ([]*Post, error) {
	rows, err := db.Query("SELECT * FROM posts ORDER BY id DESC")

	if err != nil {
		return nil, err
	}

	posts := make([]*Post, 0)

	var id int
	var title string
	var author string
	var body string
	var swritten string

	for rows.Next() {
		err := rows.Scan(&id, &title, &author, &body, &swritten)

		if err != nil {
			rows.Close()
			return nil, err
		}

		written, err := time.Parse(time.UnixDate, swritten)

		if err != nil {
			rows.Close()
			return nil, err
		}

		posts = append(posts, MakeNewPostRaw(id, title, author, body, written))
	}

	return posts, nil
}

// Getting a post based on ID
func GetPost(db *sql.DB, id int) (*Post, error) {
	var nid int
	var title string
	var author string
	var body string
	var swritten string

	err := db.QueryRow("SELECT * FROM posts WHERE id = $1", id).Scan(&nid, &title, &author, &body, &swritten)

	if err != nil {
		return nil, err
	}

	written, err := time.Parse(time.UnixDate, swritten)

	if err != nil {
		return nil, err
	}

	return &Post{
		Id:      nid,
		Title:   title,
		Author:  author,
		Body:    body,
		Written: written,
	}, nil
}

// Getting the most recent post
func MostRecent(db *sql.DB) (int, error) {
	row := db.QueryRow("SELECT id FROM posts ORDER BY id DESC")

	if row == nil {
		return 1, errors.New("Error: There are no posts.")
	}

	var id int

	err := row.Scan(&id)

	if err != nil {
		return 1, err
	}

	return id, nil
}

// Inserting a post into the database (it should be noted that the ID of
// the post is ignored, and is left to the SQL database's PRIMARY KEY
// to auto-increment)
func InsertPost(db *sql.DB, post *Post) error {
	id, err := MostRecent(db)

	exec := "INSERT INTO posts(id, title, author, body, written) values($1, $2, $3, $4, $5)"

	if err != nil {
		_, err = db.Exec(exec, 1, post.Title, post.Author, post.Body, post.Written.Format(time.UnixDate))
	} else {
		_, err = db.Exec(exec, id+1, post.Title, post.Author, post.Body, post.Written.Format(time.UnixDate))
	}

	return err
}
