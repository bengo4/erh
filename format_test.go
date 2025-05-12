package erh_test

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"

	"github.com/bengo4/erh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSourceDirectoryDepth(t *testing.T) {
	tests := []int32{1, 2, 3, erh.FileOnly, erh.FullPath}
	for i, v := range tests {
		t.Run(fmt.Sprintf("tests[%d]:%d", i, v), func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(erh.FileOnly, erh.GetSourceDirectoryDepth(), "Before Set, should be FileOnly(%d)", erh.FileOnly)
			t.Cleanup(func() {
				erh.SetSourceDirectoryDepth(erh.FileOnly)
			})
			erh.SetSourceDirectoryDepth(v)
			assert.Equal(v, erh.GetSourceDirectoryDepth(), "GetSourceDirectoryDepth should be %d", v)
		})
	}
}

func TestAddFileLine(t *testing.T) {
	tests := []struct {
		name     string
		dirDepth int32
		reWant   string
	}{
		{name: "default", reWant: `^\[testing.go:\d+\]$`},
		{name: "dirDepth==1", dirDepth: 1, reWant: `^\[testing/testing.go:\d+\]$`},
		{name: "dirDepth==2", dirDepth: 2, reWant: `^\[src/testing/testing.go:\d+\]$`},
		{name: "dirDepth==FullPath", dirDepth: erh.FullPath, reWant: `^\[/.+/src/testing/testing.go:\d+\]$`},
		{name: "dirDepth==-2", dirDepth: -2, reWant: `^\[/.+/src/testing/testing.go:\d+\]$`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NotEmpty(t, tt.reWant)

			t.Cleanup(func() {
				erh.SetSourceDirectoryDepth(erh.FileOnly)
			})
			erh.SetSourceDirectoryDepth(tt.dirDepth)

			buf := &bytes.Buffer{}
			erh.ExportAddFileLine(buf)
			assert.Regexp(t, regexp.MustCompile(tt.reWant), buf.String())
		})
	}
}

func TestFormatFileLine(t *testing.T) {
	const fullpath = "/a/b/c/d.go"
	tests := []struct {
		name     string
		file     string
		line     int
		dirDepth int32
		want     string
	}{
		{"FileOnly", fullpath, 100, erh.FileOnly, "[d.go:100]"},
		{"dirDepth==1", fullpath, 101, 1, "[c/d.go:101]"},
		{"dirDepth==2", fullpath, 102, 2, "[b/c/d.go:102]"},
		{"dirDepth==3", fullpath, 103, 3, "[a/b/c/d.go:103]"},
		{"dirDepth==4", fullpath, 104, 4, "[/a/b/c/d.go:104]"},
		{"dirDepth==100", fullpath, 105, 5, "[/a/b/c/d.go:105]"},
		{"FullPath", fullpath, 106, erh.FullPath, "[/a/b/c/d.go:106]"},
		{"dirDepth==-2", fullpath, 107, -2, "[/a/b/c/d.go:107]"},
		{"EmptyPath&FileOnly", "", 200, erh.FileOnly, "[.:200]"},
		{"EmptyPath&dirDepth==1", "", 201, erh.FileOnly, "[.:201]"},
		{"PartialPath&FileOnly", "a/b.go", 300, erh.FileOnly, "[b.go:300]"},
		{"PartialPath&dirDepth==1", "a/b.go", 301, 1, "[a/b.go:301]"},
		{"PartialPath&dirDepth==2", "a/b.go", 302, 2, "[a/b.go:302]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := erh.ExportFormatFileLine(tt.file, tt.line, tt.dirDepth)
			assert.Equal(t, tt.want, got)
		})
	}
}
