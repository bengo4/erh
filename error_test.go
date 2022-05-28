package erh_test

import (
	"fmt"
	"testing"

	"github.com/bengo4/erh"
	"github.com/stretchr/testify/assert"
)

func TestErrorf(t *testing.T) {
	tests := []struct {
		name   string
		format string
		a      []interface{}
		want   string
	}{
		{
			name:   "with a format string only",
			format: "test error", a: nil,
			want: "test error[error_test.go:33]",
		},
		{
			name:   "with a format string and an argument",
			format: "test error, p:%d", a: []interface{}{314},
			want: "test error, p:314[error_test.go:33]",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d:%s", i, tt.name), func(t *testing.T) {
			assert := assert.New(t)
			err := erh.Errorf(tt.format, tt.a...) // this is LINE:35
			assert.Equal(tt.want, err.Error())
		})
	}
}
