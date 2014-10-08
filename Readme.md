[![GoDoc][doc-img]][doc-url]
[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]

### install

```bash
go get -u github.com/coderhaoxin/go-convert
```

### example

```go
package main

import (
  "github.com/coderhaoxin/convert"
  "reflect"
  "fmt"
)

func main() {
  var s = "10086"
  i, _ := convert.Int(s)
  i32, _ := convert.Int32(s)
  i64, _ := convert.Int64(s)
  f32, _ := convert.Float32(s)
  f64, _ := convert.Float64(s)

  fmt.Println(i, i32, i64, f32, f64)
  fmt.Println(reflect.TypeOf(i), reflect.TypeOf(i32), reflect.TypeOf(i64), reflect.TypeOf(f32), reflect.TypeOf(f64))

  // json
  b := []byte(`{"name": "hello","count": 2,"users": [{"name": "hao","age": 10086},{"name": "xin","age": 12306}]}`)
  m, _ := convert.JSONmap(b)
  fmt.Println(m["name"])
  fmt.Println(m["count"])
  fmt.Println(m["users"])
}
```

### License
MIT

[doc-img]: http://img.shields.io/badge/GoDoc-reference-green.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/coderhaoxin/go-convert
[travis-img]: https://img.shields.io/travis/coderhaoxin/go-convert.svg?style=flat-square
[travis-url]: https://travis-ci.org/coderhaoxin/go-convert
[license-img]: http://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
