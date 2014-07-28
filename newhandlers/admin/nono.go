package admin

import "github.com/crockeo/personalwebsite/helpers"

func nonoHandler() (int, string) { return 200, helpers.RenderPageUnsafe("nono", struct{}{}) }
