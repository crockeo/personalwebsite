package blog

import (
	"os"
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

// Saving a Post to a file
func SavePost(post Post, path string) error {
	err := ioutil.WriteFile(path, []byte(post.Show()), os.ModeExclusive)
	return err
}
