package blog

import (
	"time"
	"strings"
	"io/ioutil"
)

type Post struct {
	Title   string
	Author  string
	Body    string
	Written time.Time
}

// Unsafely parsing the time
func unsafeParseTime(input string) time.Time {
	t, e := time.Parse(input, input)
	if e != nil {
		panic(e)
	}
	return t
}

// Creating a new time-traveled post
func NewPostWithTime(title string, author string, body string, written time.Time) Post {
	return Post{ Title  : title   ,
	             Author : author  ,
	             Body   : body    ,
	             Written: written }
}

// Creating a new Post
func NewPost(title string, author string, body string) Post {
	return NewPostWithTime(title, author, body, time.Now())
}

// Parsing out a Post
func ParsePost(input string) Post {
	lines := strings.Split(input, "\n")

	title   := ""
	author  := ""
	body    := ""
	written := time.Now()

	for index := 0; index < len(lines); index++ {
		liness := strings.SplitN(lines[index], " ", 2)

		if len(liness) == 2 {
			switch liness[0] {
			case "tit":
				title = liness[1]
			case "aut":
				author = liness[1]
			case "bod":
				body = liness[1]
			case "wri":
				written = unsafeParseTime(liness[1])
			}
		}
	}

	return NewPostWithTime(title, author, body, written)
}

// Loading a Post from a file
func LoadPost(path string) (Post, error) {
	val, err := ioutil.ReadFile(path)

	if err != nil {
		return NewPost("", "", ""), err
	}

	return ParsePost(string(val)), nil
}

// Showing a Post (converting it to a string)
func (post Post) Show() string {
	return "tit" + " " + post.Title               + "\n" +
	       "aut" + " " + post.Author              + "\n" +
	       "bod" + " " + ParseMarkdown(post.Body) + "\n" +
	       "wri" + " " + post.Written.String()
}
