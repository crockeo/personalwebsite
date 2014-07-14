package blog

import (
	"html/template"
	"strconv"
	"time"
)

// Making a Post from the title, author, and string
func MakePostRaw(id int, title string, author string, body string, written time.Time) template.HTML {
	return template.HTML(`
	<div class="col-md-8 col-md-offset-2">
		<h2>
			<a href="/blog/` + strconv.FormatInt(int64(id), 10) + `">` + title + `</a>
			<br>
			<small>By ` + author + `</small>
		</h2>

		<div class="text-justify">
			` + ParseMarkdown(body) + `
		</div>

		<h4><small>` + written.Format(time.UnixDate) + `</small></h4>
	</div>`)
}

// Making a Post with time being now
func MakePost(title string, author string, body string) template.HTML {
	return MakePostRaw(Posts(), title, author, body, time.Now())
}
