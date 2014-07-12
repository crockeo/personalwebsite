package blog

import (
	"time"
	"strconv"
	"strings"
)

type Post struct {
	Id      int
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
func NewPostWithTime(id int, title string, author string, body string, written time.Time) Post {
	return Post{ Id     : id      ,
	             Title  : title   ,
	             Author : author  ,
	             Body   : body    ,
	             Written: written }
}

// Creating a new Post
func NewPost(id int, title string, author string, body string) Post {
	return NewPostWithTime(id, title, author, body, time.Now())
}

// Parsing out a Post
func ParsePost(input string) Post {
	lines := strings.Split(input, "\n")

	id      := 0
	title   := ""
	author  := ""
	body    := ""
	written := time.Now()

	for index := 0; index < len(lines); index++ {
		liness := strings.SplitN(lines[index], " ", 2)

		if len(liness) == 2 {
			switch liness[0] {
			case "id":
				tid, err := strconv.ParseInt(liness[1], 10, 64)
				if err == nil {
					id = int(tid)
				}
			case "tit":
				title = liness[1]
			case "aut":
				author = liness[1]
			case "bod":
				body = liness[1]
			case "wri":
				twritten, err := time.Parse(liness[1], liness[1])
				if err == nil {
					written = twritten
				}
			}
		}
	}

	return NewPostWithTime(id, title, author, body, written)
}

// Showing a Post (converting it to a string)
func (post Post) Show() string {
	return "id"  + " " + strconv.FormatInt(int64(post.Id), 10) + "\n" +
	       "tit" + " " + post.Title                            + "\n" +
	       "aut" + " " + post.Author                           + "\n" +
	       "bod" + " " + ParseMarkdown(post.Body)              + "\n" +
	       "wri" + " " + post.Written.String()
}
