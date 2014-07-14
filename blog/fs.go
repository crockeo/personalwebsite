package blog

import (
	"github.com/crockeo/personalwebsite/config"
	"io/ioutil"
	"strconv"
)

// Loading a Post from a file
func LoadPostRaw(path string) (*Post, error) {
	val, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return ParsePost(string(val)), nil
}

// Loading a Post from index
func LoadPost(index int) (*Post, error) {
	return LoadPostRaw(config.PostsDir + config.PostName + strconv.FormatInt(int64(index), 10))
}

// Loading the nubmer of Posts
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

// Writing the number of Posts
func SetPosts(num int) error {
	return ioutil.WriteFile(config.PostIndexLoc, []byte(strconv.FormatInt(int64(num), 10)), 664)
}

// Incrementing the nubmer of Posts that exist
func IncPosts() error {
	return SetPosts(Posts() + 1)
}

// Loading every post
func LoadPosts() ([]*Post, error) {
	nposts := Posts()

	if nposts == 0 {
		return nil, nil
	} else {
		posts := make([]*Post, nposts)
		for i := 0; i < nposts; i++ {
			post, err := LoadPost(i)

			if err != nil {
				return nil, err
			}

			posts[i] = post
		}

		return posts, nil
	}
}
