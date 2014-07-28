package blog

import (
	"fmt"
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/database/schema"
	"github.com/crockeo/personalwebsite/helpers"
	"github.com/crockeo/personalwebsite/newhandlers/errors"
	"html/template"
)

func displayPost(post schema.Post) template.HTML {
	return template.HTML(fmt.Sprintf(`Nothing here yet`))
}

// The main handler for the blog page
func handler(db *database.DB) (int, string) {
	var posts []schema.Post
	_, err := db.Select(&posts, "SELECT * FROM posts;")

	if err != nil {
		return errors.ErrorHandler(503, err.Error())
	} else if len(posts) == 0 {
		return 200, helpers.RenderPageUnsafe("noblog", struct{}{})
	} else {
		dposts := make([]template.HTML, len(posts))
		for i, v := range posts {
			dposts[i] = displayPost(v)
		}

		return 200, helpers.RenderPageUnsafe("blog", struct {
			Posts []template.HTML
		}{
			Posts: dposts,
		})
	}
}
