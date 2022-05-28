package erh_test

import (
	"fmt"
	"testing"

	"github.com/bengo4/erh"
	"github.com/stretchr/testify/assert"
)

func TestWrapNil(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(erh.Wrap(nil))
	assert.Nil(erh.Wrap(nil, "x"))
	assert.Nil(erh.Wrap(nil, "x:%s:%d", "y", 123))
}

func TestWrap(t *testing.T) {
	assert := assert.New(t)

	e := fmt.Errorf("first")
	e0 := erh.Wrap(e)
	e1 := erh.Wrap(e0, "one")
	e2 := erh.Wrap(e1, "two, x:%s:%d", "y", 123)

	tests := []struct {
		err  error
		want string
	}{
		{e, "first"},
		{e0, "first[wrap_test.go:22]"},
		{e1, "one; first[wrap_test.go:22][wrap_test.go:23]"},
		{e2, "two, x:y:123; one; first[wrap_test.go:22][wrap_test.go:23][wrap_test.go:24]"},
	}

	for i, tt := range tests {
		assert.Equal(tt.want, tt.err.Error(), "i:%d", i)
		assert.Equal(e, erh.Cause(tt.err), "i:%d", i)
	}
}
