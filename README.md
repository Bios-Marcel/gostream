# gostream

[![GoDoc](https://godoc.org/github.com/Bios-Marcel/gostream?status.svg)](https://godoc.org/github.com/Bios-Marcel/gostream)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bios-Marcel/gostream)](https://goreportcard.com/report/github.com/Bios-Marcel/gostream)

This repository is a simple example of how streams could be implemented using golang.

There is an eager and a lazy implementation. Currently those implementations
are specifically for the `int` type and do not have any functions besides
`Filter`, `Map`, `Collect` and `FindFirst`.