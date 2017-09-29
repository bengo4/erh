package erh

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapNil(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(Wrap(nil))
	assert.Nil(Wrap(nil, "x"))
	assert.Nil(Wrap(nil, "x:%s:%d", "y", 123))
}

func TestWrap(t *testing.T) {
	assert := assert.New(t)

	e := fmt.Errorf("first")
	e0 := Wrap(e)
	e1 := Wrap(e0, "one")
	e2 := Wrap(e1, "two, x:%s:%d", "y", 123)

	tests := []struct {
		err  error
		want string
	}{
		{e, "first"},
		{e0, "first[wrap_test.go:21]"},
		{e1, "one; first[wrap_test.go:21][wrap_test.go:22]"},
		{e2, "two, x:y:123; one; first[wrap_test.go:21][wrap_test.go:22][wrap_test.go:23]"},
	}

	for i, tt := range tests {
		assert.Equal(tt.want, tt.err.Error(), "i:%d", i)
		assert.Equal(e, Cause(tt.err), "i:%d", i)
	}
}
