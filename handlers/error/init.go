package error

import "github.com/go-martini/martini"

func Init(m *martini.ClassicMartini) {
	m.NotFound(_404Handler)
}
