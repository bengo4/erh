package erh_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/bengo4/erh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorf(t *testing.T) {
	const wantLineNum = "71"
	tests := []struct {
		name     string
		dirDepth int32
		format   string
		a        []any
		reWant   string
	}{
		{
			name:   "with a format string only",
			format: "test error",
			a:      nil,
			reWant: `^test error\[error_test.go:` + wantLineNum + `\]$`,
		},
		{
			name:   "with a format string and an argument",
			format: "test error, p:%d",
			a:      []any{314},
			reWant: `^test error, p:314\[error_test.go:` + wantLineNum + `\]$`,
		},
		{
			name:     "dirDepth==1,with a format string only",
			dirDepth: 1,
			format:   "test error",
			a:        nil,
			reWant:   `^test error\[erh/error_test.go:` + wantLineNum + `\]$`,
		},
		{
			name:     "dirDepth==1,with a format string and an argument",
			dirDepth: 1,
			format:   "test error, p:%d",
			a:        []any{314},
			reWant:   `^test error, p:314\[erh/error_test.go:` + wantLineNum + `\]$`,
		},
		{
			name:     "dirDepth==FullPath,with a format string only",
			dirDepth: erh.FullPath,
			format:   "test error",
			a:        nil,
			reWant:   `^test error\[/.*/erh/error_test.go:` + wantLineNum + `\]$`,
		},
		{
			name:     "dirDepth==FullPath,with a format string and an argument",
			dirDepth: erh.FullPath,
			format:   "test error, p:%d",
			a:        []any{314},
			reWant:   `^test error, p:314\[/.*/erh/error_test.go:` + wantLineNum + `\]$`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d:%s", i, tt.name), func(t *testing.T) {
			require.NotEmpty(t, tt.reWant)
			erh.SetSourceDirectoryDepth(tt.dirDepth)
			t.Cleanup(func() {
				erh.SetSourceDirectoryDepth(erh.FileOnly)
			})

			err := erh.Errorf(tt.format, tt.a...) // LINE: 71
			assert.Regexp(t, regexp.MustCompile(tt.reWant), err.Error())
		})
	}
}
