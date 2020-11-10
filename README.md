<div align="center">
    <h1>~ colorname ~</h1>
    <strong>Find the name of colors.</strong><br><br>
    <a href="https://pkg.go.dev/github.com/zekroTJA/colorname"><img src="https://godoc.org/github.com/zekroTJA/colorname?status.svg" /></a>&nbsp;
    <a href="https://github.com/zekroTJA/colorname/actions/runs/353920630"><img src="https://github.com/zekroTJA/colorname/workflows/Main%20CI/badge.svg"/></a>&nbsp;
    <a href="https://coveralls.io/github/zekroTJA/colorname"><img src="https://coveralls.io/repos/github/zekroTJA/colorname/badge.svg" /></a>&nbsp;
    <a href="https://goreportcard.com/report/github.com/zekroTJA/colorname"><img src="https://goreportcard.com/badge/github.com/zekroTJA/colorname"/></a>&nbsp;
<br>
</div>

---

<div align="center">
    <code>go get -u github.com/zekroTJA/colorname</code>
</div>

---

## Intro

`colorname` matches a given color to the most similar looking color of the ["List of colors"](https://en.wikipedia.org/wiki/List_of_colors:_A%E2%80%93F) of Wikipedia.

This package is inspired by [jcolag/color-namer](https://github.com/jcolag/color-namer).

[Here](https://pkg.go.dev/github.com/zekroTJA/colorname) you can read the docs of this package, generated by pkg.go.dev.

---

## Usage Example

```go
package main

import (
    "fmt"

	"github.com/zekroTJA/colorname"
)

func main() {
    matches := colornames.FindNum(0x6a0dad)
    
    for _, m := range matches {
        fmt.Printf("%+v\n", m)
    }
}
```

Further examples, you can find in the [example](examples) directory.

---

Copyright (c) 2020 zekro Development (Ringo Hoffmann).  
Covered by MIT licence.