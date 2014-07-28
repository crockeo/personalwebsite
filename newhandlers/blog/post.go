package blog

import (
	"github.com/crockeo/personalwebsite/database"
	"github.com/crockeo/personalwebsite/database/schema"
	//	"github.com/crockeo/personalwebsite/helpers"
	"github.com/crockeo/personalwebsite/newhandlers/errors"
	"github.com/go-martini/martini"
	"strconv"
)

func postHandler(db *database.DB, params martini.Params) (int, string) {
	id, err := strconv.ParseInt(params["id"], 64, 10)

	if err != nil {
		return errors.ErrorHandler(404, "Post does not exist.")
	}

	var post schema.Post
	err = db.SelectOne(&post, "SELECT * FROM posts WHERE id = ?", id)

	if err != nil {
		return errors.ErrorHandler(404, err.Error())
	}

	return 200, "Nothing to see here"
}
