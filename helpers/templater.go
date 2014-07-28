package helpers

import (
	"html/template"
	"io"
	"log"
)

const (
	HeaderName = "header"
	FooterName = "footer"
)

// The Page type (for injecting data)
type Page struct {
	Name string
	Data interface{}
}

// Loading a template safely (just by name)
func LoadTemplate(name string) (*template.Template, error) {
	temp, err := template.ParseFiles("templates/" + name + ".html")

	return temp, err
}

// Loading a template unsafely (just by name)
func LoadTemplateUnsafe(name string) *template.Template {
	temp, err := LoadTemplate(name)

	if err != nil {
		log.Fatal(err)
	}

	return temp
}

// Loading the header and footer templates
func loadHeader() *template.Template { return LoadTemplateUnsafe(HeaderName) }
func loadFooter() *template.Template { return LoadTemplateUnsafe(FooterName) }

// Sending a file with the header and
// footer templates included
func SendPage(w io.Writer, name string, data interface{}) error {
	page, err := LoadTemplate(name)

	if err != nil {
		return err
	}

	header := loadHeader()
	footer := loadFooter()

	v := Page{Name: name,
		Data: data}

	header.Execute(w, v)
	page.Execute(w, v)
	footer.Execute(w, v)

	return nil
}

// Loading a template to a string
func RenderPage(name string, data interface{}) (string, error) {
	sw := NewStringWriter()

	err := SendPage(sw, name, data)

	if err != nil {
		return "", err
	}

	return sw.String(), nil
}

// Unsafely rendering the page
func RenderPageUnsafe(name string, data interface{}) string {
	str, err := RenderPage(name, data)
	if err != nil {
		log.Fatal(err)
	}
	return str
}
