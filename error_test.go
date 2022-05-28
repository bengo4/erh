package erh_test

import (
	"testing"

	"github.com/bengo4/erh"
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
			"test error[error_test.go:31]",
		},
		{
			"test error, p:%d",
			[]interface{}{314},
			"test error, p:314[error_test.go:31]",
		},
	}

	for i, tt := range tests {
		err := erh.Errorf(tt.f, tt.a...) // this is LINE:31
		assert.Equal(tt.want, err.Error(), "i:%d", i)
	}
}
