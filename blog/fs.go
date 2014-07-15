package blog

import (
	"github.com/crockeo/personalwebsite/config"
	"html/template"
	"io/ioutil"
	"strconv"
)

// Loading a Post from a file
func LoadPostRaw(path string) (template.HTML, error) {
	val, err := ioutil.ReadFile(path)

	if err != nil {
		return template.HTML(""), err
	}

	return template.HTML(val), nil
}

// Saving a Post to a file
func SavePostRaw(path string, post template.HTML) error {
	return ioutil.WriteFile(path, []byte(post), 0664)
}

// Loading a Post from an index
func LoadPost(index int) (template.HTML, error) {
	return LoadPostRaw(config.PostsDir + config.PostName + strconv.FormatInt(int64(index), 10))
}

// Saving a Post to an index
func SavePost(index int, post template.HTML) error {
	return SavePostRaw(config.PostsDir+config.PostName+strconv.FormatInt(int64(index), 10), post)
}

// Getting the number of Posts that exist
func Posts() int {
	val, err := ioutil.ReadFile(config.PostIndexLoc)

	if err != nil {
		return 0
	}

	ret, err := strconv.ParseInt(string(val), 10, 64)

	if err != nil {
		return 0
	}

	return int(ret)
}

// Setting the number of Posts
func SetPosts(num int) error {
	return ioutil.WriteFile(config.PostIndexLoc, []byte(strconv.FormatInt(int64(num), 10)), 0644)
}

// Saving the next Post
func SavePostNext(post template.HTML) error {
	posts := Posts()
	err := SavePost(posts, post)
	if err != nil {
		return err
	}

	return SetPosts(posts + 1)
}

// Loading every post
func LoadPosts() ([]template.HTML, error) {
	nposts := Posts()

	if nposts == 0 {
		return nil, nil
	} else {
		posts := make([]template.HTML, nposts)
		indx := nposts - 1
		for i := 0; i < nposts; i++ {
			post, err := LoadPost(i)

			if err != nil {
				return nil, err
			}

			posts[indx] = post
			indx--
		}

		return posts, nil
	}
}
