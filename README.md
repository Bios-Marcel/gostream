# gostream

[![GoDoc](https://godoc.org/github.com/Bios-Marcel/gostream?status.svg)](https://godoc.org/github.com/Bios-Marcel/gostream)
[![builds.sr.ht status](https://builds.sr.ht/~biosmarcel/gostream/arch.yml.svg)](https://builds.sr.ht/~biosmarcel/gostream/arch.yml?)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bios-Marcel/gostream)](https://goreportcard.com/report/github.com/Bios-Marcel/gostream)

This repository contains a generic implementation for streams, similar to the
stream API that Java has. For a list of available functions, check the
[documentation](https://godoc.org/github.com/Bios-Marcel/gostream#GenericStreamEntityStream).

## Dude, what are those stream things you are talking about

I am glad you asked! Streams is basically a small concept that allows
you to process data in a declerative way. The produced code is mostly easy
to understand and can theoretically automatically be scheduled onto multiple
threads, or in the case of Go, coroutines.

Overall streams consist of two types of methods, those are terminating methods
and non-terminating methods. A terminating method ends the call-chain, as it
returns some kind of end result. A non-terminating methods simply returns the
stream itself. A stream cannot be used multiple times, every time you want to
use a stream, you have to create a new one.

## Usage

This library is implemented using [genny](https://github.com/cheekybits/genny),
that means that each needed implementation has to be generated before starting
to compile your actual code.

1. You'll have to install `genny` in order to generate the necessary source
    ```shell
    go install github.com/cheekybits/genny
    ```
2. Pull the repository
    ```shell
    go get github.com/Bios-Marcel/gostream
    ```
3. Generate the versions you need
    ```shell
    genny -in="$GOPATH/src/github.com/Bios-Marcel/gostream/stream.go" /
        -out="folder/outputfile.go" -pkg="newpackagename" /
        "GenericStreamEntity=desiredtypes"
4. Compile your code

## Examples

Get all even numbers, multiply them by two and sum up the leftover values.

For this example I am just going to use an eager stream.

```go
data := []int{1,2,3,4,5,6,7,8,9,10}
summedEvens := gostream.
    StreamIntEager(data).
    Filter(func(value int) bool {return value%2 == 0}).
    Map(func(value int) int {return value * 2})
    Reduce(func(one, two int) int {return one + two})

fmt.Println(summedEvens)
```

However, the usage of laziness is way more interesting, for example
look at this code:

```go
testData := []int{1, 2, 3, 4, 5}
firstValueValid := gostream.
    StreamIntLazy(testData).
    Filter(func(value int) bool { return value != 2 }).
    Map(func(value int) int { return value * 4 }).
    FindFirst()
```

The functions passed to `Filter(...)` and `Map(...)` will be executed exactly
once, since `FindFirst()` will stop executing after it finds any value.

In an eager stream, the function passed to `Filter(...)` would execute five
times and the function passed to `Map(...)` would execute four times.

In case you don't care wether the implementation should be eager or lazy,
simply use the method `gostream.StreamGenericStreamEntity([]int) IntStream`.

## Making use of parallelism

In order to get the maximum out of Go and streams, having streams make use of
parallelism would be very useful.

A stream should only be parallel if the user ask the framework to do so,
otherwise a lot of logic is required in order to find out if parallelism would
make sense or not. Introducing such logic would create an overhead and make
the code more prone to errors. Therefore, a parallel stream will always create
multiple channels per step, no matter if the terminal action is a `Collect()`
or a `FindFirst()`.

A downside would be that in case of a `FindFirst()` the stream might execute
more code than necessary, since the go routines aren't killable from outisde,
unless the user prepares them to be killable, but that would introduce too
much boilerplate code when using the parallel api. However, since the code
executed shouldn't have side effects, that might be fine.