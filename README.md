# glog

**[glog](https://pkg.go.dev/github.com/winterssy/glog)** provides a simple logger for Go.

![Build](https://img.shields.io/github/workflow/status/winterssy/glog/Test/master?logo=appveyor) [![Go Report Card](https://goreportcard.com/badge/github.com/winterssy/glog)](https://goreportcard.com/report/github.com/winterssy/glog) [![GoDoc](https://img.shields.io/badge/godoc-reference-5875b0)](https://pkg.go.dev/github.com/winterssy/glog) [![License](https://img.shields.io/github/license/winterssy/glog.svg)](LICENSE)

## Install

```sh
go get -u github.com/winterssy/glog
```

## Usage

```go
import "github.com/winterssy/glog"
```

## Quick Start

```go
package main

import "github.com/winterssy/glog"

func init() {
	glog.ReplaceGlobal(glog.New(os.Stderr, "", glog.LstdFlags, glog.Ldebug))
}

func main() {
	glog.Debug("hello world")
	// 2020-01-01 00:00:00,000 [DEBUG] hello world
}
```

## License

**[MIT](LICENSE)**
