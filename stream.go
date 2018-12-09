package gostream

//IntStream is a stream over an int array. Depending on its implemenation
//its behaviour may vary. The default implementations are an eager stream
//and a lazy stream.
type IntStream interface {
	//Filter only keeps the values that meet the filter condition.
	Filter(func(value int) bool) IntStream
	//Map changes the values according to the mapping function.
	Map(func(value int) int) IntStream
	//Reduces combines all values using the given reduce function.
	Reduce(func(one, two int) int) *int
	//FindFirst returns the first element found.
	FindFirst() *int
	//Collect returns a copy of the internally used int array.
	Collect() []int
}

//StreamInts creates an IntStream using the lazy implementation of IntStream.
func StreamInts(data []int) IntStream {
	return StreamIntsLazy(data)
}

func reduceIntArray(reduceFunction func(valueOne, valueTwo int) int, data []int) *int {
	lengthOfData := len(data)

	if lengthOfData == 0 {
		return nil
	}

	if lengthOfData == 1 {
		return &data[0]
	}

	var value int
	if lengthOfData == 2 {
		value = reduceFunction(data[0], data[1])
	} else {
		value = data[0]
		for index := 1; index < lengthOfData; index++ {
			value = reduceFunction(value, data[index])
		}
	}

	return &value
}
