package blog

import "github.com/russross/blackfriday"

// Parsing markdown into an HTML-ey file
func ParseMarkdown(input string) string {
	return string(blackfriday.MarkdownCommon([]byte(input)))
}
