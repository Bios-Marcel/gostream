package gostream

//IntStream is a stream over an int array. Depending on its implemenation
//its behavour may vary. The default implementations are an eager stream
//and a lazy stream.
type IntStream interface {
	//Filter only keeps the values that meet the filter condition.
	Filter(func(value int) bool) IntStream
	//Map changes the values according to the mapping function.
	Map(func(value int) int) IntStream
	//FindFirst returns the first element found.
	FindFirst() *int
	//Collect returns a copy of the internally used int array.
	Collect() []int
}
