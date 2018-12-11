package gostream_test

//go:generate genny -in=stream.go -out=intstream.go gen "GenericStreamEntity=int"

import (
	"testing"

	"github.com/Bios-Marcel/gostream"
)

func TestStreamEagerCollect(t *testing.T) {
	testData := []int{1, 2, 3}
	resultData := gostream.
		StreamIntEager(testData).
		Filter(func(value int) bool { return value != 2 }).
		Map(func(value int) int { return value * 4 }).
		Collect()

	valueAtIndexZero := resultData[0]
	if valueAtIndexZero != 4 {
		t.Errorf("Error, value at index 0 was %d, but should have been 4.", valueAtIndexZero)
	}

	valueAtIndexOne := resultData[1]
	if valueAtIndexOne != 12 {
		t.Errorf("Error, value at index 1 was %d, but should have been 12.", valueAtIndexOne)
	}
}

func TestStreamEagerFindFirst(t *testing.T) {
	testData := []int{1, 2, 3}
	firstValueValid := gostream.
		StreamIntEager(testData).
		Filter(func(value int) bool { return value%2 == 1 }).
		Map(func(value int) int { return value * 4 }).
		FindFirst()

	if *firstValueValid != 4 {
		t.Errorf("Error, first value was %d, but should have been 4.", firstValueValid)
	}

	firstValueInvalid := gostream.
		StreamIntEager(testData).
		Filter(func(value int) bool { return value == -1 }).
		FindFirst()

	if firstValueInvalid != nil {
		t.Errorf("Error, first value was %d, but should have been nil.", firstValueInvalid)
	}
}

func TestStreamEagerReduce(t *testing.T) {
	testData := []int{1, 2, 3}
	reducedValue := gostream.
		StreamIntEager(testData).
		Filter(func(value int) bool { return value != 1 }).
		Map(func(value int) int { return value * 4 }).
		Reduce(func(one, two int) int { return one + two })

	if *reducedValue != 20 {
		t.Errorf("Error, first value was %d, but should have been 20.", reducedValue)
	}

	reducedValue = gostream.
		StreamIntEager(testData).
		Filter(func(value int) bool { return value == -1 }).
		Reduce(func(one, two int) int { return one + two })

	if reducedValue != nil {
		t.Errorf("Error, first value was %d, but should have been nil.", reducedValue)
	}
}

func TestStreamEagerness(t *testing.T) {
	testData := []int{1, 2, 3}
	mapIterationCounter := 0
	gostream.
		StreamIntEager(testData).
		Filter(func(value int) bool { return value%2 == 1 }).
		Map(func(value int) int {
			mapIterationCounter++
			return value * 4
		}).
		FindFirst()

	if mapIterationCounter != 2 {
		t.Errorf("Error, map should have been called 2 times, but was called %d times.", mapIterationCounter)
	}
}
