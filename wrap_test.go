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
	e0 := fmt.Errorf("first")
	e1 := erh.Wrap(e0)                                             // LINE: 20
	e2 := erh.Wrap(e1, "simple message")                           // LINE: 21
	e3 := erh.Wrap(e2, "formatted with params, x:%s:%d", "y", 123) // LINE: 22

	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "not wrapped",
			err:  e0,
			want: "first",
		},
		{
			name: "wrapped 1x",
			err:  e1,
			want: "first[wrap_test.go:20]",
		},
		{
			name: "wrapped 2x",
			err:  e2,
			want: "simple message; first[wrap_test.go:20][wrap_test.go:21]",
		},
		{
			name: "wrapped 3x",
			err:  e3,
			want: "formatted with params, x:y:123; simple message; first[wrap_test.go:20][wrap_test.go:21][wrap_test.go:22]",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d:%s", i, tt.name), func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.want, tt.err.Error(), "i:%d", i)
			assert.Equal(e0, erh.Cause(tt.err), "i:%d", i)
		})
	}
}
