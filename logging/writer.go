package logging

import "io"

type defaultLogWriter struct{}

func (lw *defaultLogWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

//NewDefaultLogWriter return default log writer
func NewDefaultLogWriter() io.Writer {
	return &defaultLogWriter{}
}
