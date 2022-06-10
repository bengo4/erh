package erh_test

import (
	"errors"
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
	e1 := erh.Wrap(e0)                                             // LINE: 21
	e2 := erh.Wrap(e1, "simple message")                           // LINE: 22
	e3 := erh.Wrap(e2, "formatted with params, x:%s:%d", "y", 123) // LINE: 23

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
			want: "first[wrap_test.go:21]",
		},
		{
			name: "wrapped 2x",
			err:  e2,
			want: "simple message; first[wrap_test.go:21][wrap_test.go:22]",
		},
		{
			name: "wrapped 3x",
			err:  e3,
			want: "formatted with params, x:y:123; simple message; first[wrap_test.go:21][wrap_test.go:22][wrap_test.go:23]",
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

func TestWrap_Is(t *testing.T) {
	e0 := fmt.Errorf("first")
	e1 := erh.Wrap(e0)
	e2 := erh.Wrap(e1, "simple message")
	errs := map[string]error{
		"e0": e0,
		"e1": e1,
		"e2": e2,
	}

	tests := []struct {
		err    string
		target string
		want   bool
	}{
		{err: "e0", target: "e0", want: true},
		{err: "e0", target: "e1", want: false},
		{err: "e0", target: "e2", want: false},
		{err: "e1", target: "e0", want: true},
		{err: "e1", target: "e1", want: true},
		{err: "e1", target: "e2", want: false},
		{err: "e2", target: "e0", want: true},
		{err: "e2", target: "e1", want: true},
		{err: "e2", target: "e2", want: true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d:Is(%s,%s)", i, tt.err, tt.target), func(t *testing.T) {
			assert := assert.New(t)
			err := errs[tt.err]
			if !assert.NotNil(err) {
				return
			}
			target := errs[tt.target]
			if !assert.NotNil(target) {
				return
			}

			assert.Equal(tt.want, errors.Is(err, target))
		})
	}
}
