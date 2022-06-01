# Series

![test](https://github.com/WinPooh32/series/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/WinPooh32/series)](https://goreportcard.com/report/github.com/WinPooh32/series)
[![Go Reference](https://pkg.go.dev/badge/github.com/WinPooh32/series.svg)](https://pkg.go.dev/github.com/WinPooh32/series)

Liteweight library of series data processing functions written in pure ideomatic Go. Inspired by python library [Pandas](https://github.com/pandas-dev/pandas).

For enabling **float32** use build tag `series_f32`, otherwise **float64** will be used as data type.

Build command example:

``` shell
go build -tags series_f32
```

[Code Examples](https://github.com/WinPooh32/fta/blob/master/fta.go)
