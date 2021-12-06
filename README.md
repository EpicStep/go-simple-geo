go-simple-geo is a library for simple geo calculations.

[![Go Reference](https://pkg.go.dev/badge/github.com/EpicStep/go-simple-geo.svg)](https://pkg.go.dev/github.com/EpicStep/go-simple-geo)
[![Lint](https://github.com/EpicStep/go-simple-geo/actions/workflows/lint.yml/badge.svg)](https://github.com/EpicStep/go-simple-geo/actions/workflows/lint.yml)
[![Tests](https://github.com/EpicStep/go-simple-geo/actions/workflows/go.yml/badge.svg)](https://github.com/EpicStep/go-simple-geo/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/EpicStep/go-simple-geo/branch/master/graph/badge.svg?token=UE3A8O81TA)](https://codecov.io/gh/EpicStep/go-simple-geo)
[![Go Report Card](https://goreportcard.com/badge/github.com/EpicStep/go-simple-geo)](https://goreportcard.com/report/github.com/EpicStep/go-simple-geo)

----

## Installation
```bash
go get github.com/EpicStep/go-simple-geo
```

## Example

```go
package main

import (
	"github.com/EpicStep/go-simple-geo"
	"log"
	"math"
)

func main() {
	origin := geo.Vec3d{
		X: 0,
		Y: 0,
		Z: 0,
	}

	c1 := geo.Coordinates{
		Latitude:  math.Pi,
		Longitude: 0,
	}

	log.Println(origin.SquareDistance(c1.ToVec3d()))
}
```