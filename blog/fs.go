package blog

import (
	"github.com/crockeo/personalwebsite/config"
	"io/ioutil"
	"os"
	"strings"
)

// Loading a Post from a file
func LoadPost(path string) (*Post, error) {
	val, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return ParsePost(string(val)), nil
}

// Loading a splice of Posts
func LoadPosts(path string) ([]*Post, error) {
	val, err := ioutil.ReadFile(path)

	if err != nil {
		return []*Post{}, err
	}

	uposts := strings.Split(string(val), "spl\n")
	posts := make([]*Post, len(uposts))

	for i := 0; i < len(uposts); i++ {
		posts[i] = ParsePost(uposts[i])
	}

	return posts, nil
}

// Loading the default Posts
func LoadDefaultPosts() ([]*Post, error) {
	return LoadPosts(config.PostsLoc)
}

// Creating a Posts file
func CreatePostsFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ioutil.WriteFile(path, []byte(""), os.ModeExclusive)
	}

	return nil
}

// Creating the default Posts file
func CreateDefaultPostsFile() error { return CreatePostsFile(config.PostsLoc) }

// Appending a Posts file
func AppendPostsFile(path string, post Post) error {
	err := CreatePostsFile(path)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(post.String()), os.ModeAppend)

	return err
}

// Appending the default Posts file
func AppendDefaultPostsFile(post Post) error { return AppendPostsFile(config.PostsLoc, post) }
