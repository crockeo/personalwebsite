package home

import "github.com/crockeo/personalwebsite/helpers"

func handler() string { return helpers.RenderPageUnsafe("home", struct{}{}) }
