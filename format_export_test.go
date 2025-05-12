package erh

// ExportAddFileLine is a test helper to call addFileLine from outside the package.
// It is defined as a variable (not a function) not to add a stack frame,
// so that runtime.Caller(2) in addFileLine refers to the correct caller in tests.
var ExportAddFileLine = addFileLine

// ExportFormatFileLine is a test helper to call formatFileLine from outside the package.
func ExportFormatFileLine(file string, line int, dirDepth int32) string {
	return formatFileLine(file, line, dirDepth)
}
