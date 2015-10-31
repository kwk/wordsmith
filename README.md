# Wordsmith

*The tool is not production ready and does not have test coverage. It developed
as a showcase for [my recent blog post
article](http://blog.ralch.com/tutorial/golang-code-generation-tool-implementation/).*

A `go:generate` command line tool that produces
[io.WriterTo](https://golang.org/pkg/io/#WriterTo) implementation:

```
type WriterAt interface {
        WriteAt(p []byte, off int64) (n int, err error)
}
```

#### Installation

To install the package and command line tool, use the following:

```
$ go get github.com/svett/wordsmith
```

#### Usage

The `wordsmith` command should be installed in your `$GOPATH\bin` directory
and can be invoked in the following format:

```
wordsmith -pointer -type=<type> -format=json
```

##### Available parameters:

 - `format` string encoding format (default "json")
 - `pointer` determines whether a type is a pointer or not
 - `type` string type that hosts io.WriterTo interface implementation
 - `package` string package name that hosts the type

#### Getting started

The command line tool can be use as `go:generate` subcommand. The following code
snippet demonstrates that:

```
package geometry

//go:generate wordsmith -type=Point -format=json
type Point struct {
  X float32
    Y float32
}
```

Then `go:generate` command should be executed:

```
$ go generate
```

The `wordsmith` tool generates `point_writerto.go` file in the same package,
where `Point` struct lives:

```
// @filaname: point_writerto.go
package geometry

import (
    "encoding/json"
    "io"
    )

func (obj Point) WriteTo(writer io.Writer) (int64, error) {
  data, err := json.Marshal(&obj)
    if err != nil {
      return 0, err
    }
  length, err := writer.Write(data)
    return int64(length), err
}
```

*Note that right now it supports only `json` format.*

#### Author

[Svett Ralchev](http://www.ralch.com)

#### License

[MIT License (MIT)](https://opensource.org/licenses/MIT)
