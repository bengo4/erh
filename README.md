# erh
Simple Error Handling Functions for Go.

```
go get github.com/bengo4/erh@latest

import "github.com/bengo4/erh"
```

When an `error` returned from a nested function, it may be difficult for you to find the birth place or the pathway of the `error`.

To solve this problem, the `erh` package provides functions below, similar to [github.com/pkg/errors](https://github.com/pkg/errors):

* `Wrap`
* `Cause`
* `Errorf`

Also `erh` provides function `As`, a wrapper of [errors.As](https://pkg.go.dev/errors#As).

The directory depth of the source file recorded in error messages can be changed by `SetSourceDirectoryDepth`.

## Wrap

The `erh.Wrap` function creates a new error based on the original error. 

You can add the short filename and the line number of the place where `erh.Wrap` is used to the original error message.

```
func main() {
	_, err := ioutil.ReadFile("file-not-exists")
	fmt.Println(erh.Wrap(err))
}
```

result:

```
open file-not-exists: no such file or directory[main.go:12]
```

It is also possible to add an extra message like below.

```
func main() {
	_, err := ioutil.ReadFile("file-not-exists")
	fmt.Println(erh.Wrap(err, "read failed"))
}
```

result:

```
read failed; open file-not-exists: no such file or directory[main.go:12]
```

You can also use `fmt.Sprintf` style to add a message.

```
func main() {
	_, err := ioutil.ReadFile("file-not-exists")
	fmt.Println(erh.Wrap(err, "read failed, version:%d", 1))
}
```

result:

```
read failed, version:1; open file-not-exists: no such file or directory[main.go:12]
```

If the original error is `nil`, `erh.Wrap(nil)` returns `nil`.

The error wrapped by Wrap can be retrieved by errors.Unwrap. So errors.Is can be used with Wrap.

The very first error of repeatedly wrapped errors can be retrieved by Cause.

## Cause

The `erh.Cause` function regains the original error from the one wrapped by `erh.Wrap`.

```
func main() {
	_, err := ioutil.ReadFile("file-not-exists")
	errW := erh.Wrap(err)
	fmt.Println(errW)
	fmt.Println(erh.Cause(errW))
}
```

result:

```
open file-not-exists: no such file or directory[main.go:12]
open file-not-exists: no such file or directory
```

## Errorf

With the `erh.Errorf`, you can create a new error with a message containing the short filename and the line number of the place where `erh.Errorf` is used.

```
func main() {
	fmt.Println(erh.Errorf("something wrong, version:%d", 1))
}
```

results:
```
something wrong, version:1[main.go:10]
```

## As

The function `erh.As` is a wrapper of [errors.As](https://pkg.go.dev/errors#As).

An example of using `erh.As` is below.

```go
target, ok := erh.As[*fs.PathError](err)
```

This is equivalent to the following code.

```go
var (
    target *fs.PathError
    _      error = target
)
ok := errors.As(err, &target)
```

## SetSourceDirectoryDepth

By default, erh records only the base names of the source files in error messages. You can include additional trailing directory levels by specifying a value with `SetSourceDirectoryDepth`. If you set `erh.FullPath` or a negative value, the full file path will be recorded.

```
func init() {
	erh.SetSourceDirectoryDepth(1)
}

func main() {
	fmt.Println(erh.Errorf("error happened"))
}
```

results:
```
error happend[some_directory/main.go:10]
```
