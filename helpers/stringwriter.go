package helpers

// The StringWriter type
type StringWriter struct {
	Bytes []byte
}

// Making a new StringWriter
func NewStringWriter() *StringWriter {
	return &StringWriter{
		Bytes: make([]byte, 0),
	}
}

// Writing bytes to the string writer
func (this *StringWriter) Write(p []byte) (int, error) {
	this.Bytes = append(this.Bytes, p...)
	return len(p), nil
}

// Converting the StringWriter's bytes to a string
func (this *StringWriter) String() string { return string(this.Bytes) }
