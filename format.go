package erh

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
)

func addFileLine(w io.Writer) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		panic("no caller")
	}
	_, err := fmt.Fprintf(w, "[%s:%d]", filepath.Base(file), line)
	if err != nil {
		panic(err)
	}
}
