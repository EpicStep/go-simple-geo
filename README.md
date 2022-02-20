go-simple-geo is a library for simple geo calculations.

[![Go Reference](https://pkg.go.dev/badge/github.com/EpicStep/go-simple-geo.svg)](https://pkg.go.dev/github.com/EpicStep/go-simple-geo/v2)
[![Lint](https://github.com/EpicStep/go-simple-geo/actions/workflows/lint.yml/badge.svg)](https://github.com/EpicStep/go-simple-geo/actions/workflows/lint.yml)
[![Tests](https://github.com/EpicStep/go-simple-geo/actions/workflows/go.yml/badge.svg)](https://github.com/EpicStep/go-simple-geo/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/EpicStep/go-simple-geo/branch/master/graph/badge.svg?token=UE3A8O81TA)](https://codecov.io/gh/EpicStep/go-simple-geo)
[![Go Report Card](https://goreportcard.com/badge/github.com/EpicStep/go-simple-geo)](https://goreportcard.com/report/github.com/EpicStep/go-simple-geo)

----

## Installation
```bash
go get github.com/EpicStep/go-simple-geo/v2
```

## Example

```go
package main

import (
	"fmt"
	
	"github.com/EpicStep/go-simple-geo/geo/v2"
)

func main() {
	c1 := geo.NewCoordinatesFromDegrees(55.7522, 37.6156)
	c2 := geo.NewCoordinatesFromDegrees(59.9386, 30.3141)

	fmt.Println(c1.Distance(c2))
}
```