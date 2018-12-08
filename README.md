# gostream

[![GoDoc](https://godoc.org/github.com/Bios-Marcel/gostream?status.svg)](https://godoc.org/github.com/Bios-Marcel/gostream)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bios-Marcel/gostream)](https://goreportcard.com/report/github.com/Bios-Marcel/gostream)

This repository is a simple example of how streams could be implemented using golang.

There is an eager and a lazy implementation. Currently those implementations
are specifically for the `int` type and do not have any functions besides
`Filter`, `Map`, `Collect`, `Reduce` and `FindFirst`.

## Examples

Get all even numbers, multiply them by two and sum up the leftover values.

For this example I am just going to use an eager stream.

```go
data := []int{1,2,3,4,5,6,7,8,9,10}
summedEvens := gostream.
    StreamIntsEager(data).
    Filter(func(value int) bool {return value % 2 == 0}).
    Map(func(value int) int {return value * 2})
    Reduce(func(one, two int) int {return one + two})

fmt.Println(summedEvens)
```