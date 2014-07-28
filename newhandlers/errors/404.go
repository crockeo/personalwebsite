package errors

func _404Handler() (int, string) { return ErrorHandler(404, "Page not found!") }
