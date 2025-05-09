package erh

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
)

// Constants to be used with SetSourceDirectoryDepth.
const (
	FileOnly int32 = 0  // only the file name is included in error messages
	FullPath int32 = -1 // the full file path is included in error messages
)

// sourceDirectoryDepth controls the depth of directories included in the path for error messages.
var sourceDirectoryDepth int32 = FileOnly

// SetSourceDirectoryDepth sets the directory depth included in error messages.
//
// The directory depth of the source file recorded in error messages can be changed by this function.
// This controls how many trailing directories are included in the file path for error messages.
// If set to FileOnly(0), only the base name is included. If set to FullPath(-1) or a negative value, the full path is included.
func SetSourceDirectoryDepth(depth int32) {
	atomic.StoreInt32(&sourceDirectoryDepth, depth)
}

// GetSourceDirectoryDepth returns the current directory depth included in error messages.
func GetSourceDirectoryDepth() int32 {
	return atomic.LoadInt32(&sourceDirectoryDepth)
}

// addFileLine writes the file and line information to the given writer for error messages.
// This function records the file and line number 2 frames up the call stack,
// so it should only be called from public functions where recording the caller's location is necessary.
func addFileLine(w io.Writer) {
	_, file, line, _ := runtime.Caller(2)
	dirDepth := GetSourceDirectoryDepth()
	if _, err := fmt.Fprint(w, formatFileLine(file, line, dirDepth)); err != nil {
		log.Printf("erh.addFileLine error: %v", err)
	}
}

// formatFileLine returns a formatted string for error messages, including the file path (with the specified directory depth) and line number.
func formatFileLine(file string, line int, directoryDepth int32) string {
	if directoryDepth == 0 {
		file = filepath.Base(file)
	} else if directoryDepth > 0 {
		parts := strings.Split(filepath.ToSlash(file), "/")
		if len(parts) > int(directoryDepth)+1 {
			file = filepath.Join(parts[len(parts)-int(directoryDepth)-1:]...)
		}
	}
	return fmt.Sprintf("[%s:%d]", file, line)
}
