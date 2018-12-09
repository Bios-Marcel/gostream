# gostream

[![GoDoc](https://godoc.org/github.com/Bios-Marcel/gostream?status.svg)](https://godoc.org/github.com/Bios-Marcel/gostream)
[![builds.sr.ht status](https://builds.sr.ht/~biosmarcel/gostream/arch.yml.svg)](https://builds.sr.ht/~biosmarcel/gostream/arch.yml?)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bios-Marcel/gostream)](https://goreportcard.com/report/github.com/Bios-Marcel/gostream)

This repository is a simple example of how streams could be implemented using golang.

There is an eager and a lazy implementation. Currently those implementations
are specifically for the `int` type and do not have any functions besides
`Filter`, `Map`, `Collect`, `Reduce` and `FindFirst`.

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

## Why only for integers

Well, as you may have noticed, Go doesn't have generics and therefore it isn't
possible to write something like this in a generic way, unless you are willing
to fill your code with a lot of `interface{}` and casts.

In the future, a look at [genny](https://github.com/cheekybits/genny) might be
worth considering.

I might just implement streams for all the primitive datatypes and for
`interface`. Anyhow, this project is first of all just something I do for
fun, so don't expect it to have proper support and such.

## Examples

Get all even numbers, multiply them by two and sum up the leftover values.

For this example I am just going to use an eager stream.

```go
data := []int{1,2,3,4,5,6,7,8,9,10}
summedEvens := gostream.
    StreamIntsEager(data).
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
    StreamIntsLazy(testData).
    Filter(func(value int) bool { return value != 2 }).
    Map(func(value int) int { return value * 4 }).
    FindFirst()
```

The functions passed to `Filter(...)` and `Map(...)` will be executed exactly
once, since `FindFirst()` will stop executing after it finds any value.

In an eager stream, the function passed to `Filter(...)` would execute five
times and the function passed to `Map(...)` would execute four times.

In case you don't care wether the implementation should be eager or lazy,
simply use the method `gostream.StreamInts([]int) IntStream`.

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