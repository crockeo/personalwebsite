package blog

import (
	"os"
	"strings"
	"io/ioutil"
)

// Loading a Post from a file
func LoadPost(path string) (Post, error) {
	val, err := ioutil.ReadFile(path)

	if err != nil {
		return NewPost(0, "", "", ""), err
	}

	return ParsePost(string(val)), nil
}

// Loading a splice of Posts
func LoadPosts(path string) ([]Post, error) {
	val, err := ioutil.ReadFile(path)

	if err != nil {
		return []Post{}, err
	}

	uposts := strings.Split(string(val), "spl\n")
	posts  := make([]Post, len(uposts))

	for i := 0; i < len(uposts); i++ {
		posts[i] = ParsePost(uposts[i])
	}

	return posts, nil
}

// Saving a Post to a file
func SavePost(post Post, path string) error {
	err := ioutil.WriteFile(path, []byte(post.Show()), os.ModeExclusive)
	return err
}
