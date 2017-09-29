package erh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorf(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		f    string
		a    []interface{}
		want string
	}{
		{
			"test error",
			nil,
			"test error[error_test.go:30]",
		},
		{
			"test error, p:%d",
			[]interface{}{314},
			"test error, p:314[error_test.go:30]",
		},
	}

	for i, tt := range tests {
		err := Errorf(tt.f, tt.a...) // this is LINE:30
		assert.Equal(tt.want, err.Error(), "i:%d", i)
	}
}
