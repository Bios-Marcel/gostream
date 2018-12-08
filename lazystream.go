package gostream

type lazyIntStream struct {
	data      []int
	functions []interface{}
}

type intFilter func(value int) bool
type intMapper func(value int) int

//StreamIntsLazy creates a lazy IntStream that uses a copy of the passed array.
func StreamIntsLazy(data []int) IntStream {
	defensiveCopy := make([]int, len(data))
	copy(defensiveCopy, data)
	return &lazyIntStream{
		data: defensiveCopy,
	}
}

func (intStream *lazyIntStream) Filter(filterFunction func(value int) bool) IntStream {
	intStream.functions = append(intStream.functions, (intFilter)(filterFunction))
	return intStream
}

func (intStream *lazyIntStream) Map(mapFunction func(value int) int) IntStream {
	intStream.functions = append(intStream.functions, (intMapper)(mapFunction))
	return intStream
}

func (intStream *lazyIntStream) FindFirst() *int {
FINDFIRST_VALUE_LOOP:
	for _, value := range intStream.data {
		for _, function := range intStream.functions {
			castFilter, ok := function.(intFilter)
			if ok {
				if castFilter(value) {
					//Continue, since value fits the filter
					continue
				} else {
					//Value has to be sorted out, therefore skip to next value
					continue FINDFIRST_VALUE_LOOP
				}
			}

			castMapper, ok := function.(intMapper)
			if ok {
				value = castMapper(value)
				continue
			}

		}

		//As this is a terminating function, quit if we successfully end a single iteration.
		return &value
	}

	return nil
}

func (intStream *lazyIntStream) Reduce(reduceFunction func(one, two int) int) *int {
	//This implementation is just being lazy, instead of a proper optimized solution,
	//I'll simply call collect and reduce the result.
	return reduceIntArray(reduceFunction, intStream.Collect())
}

func (intStream *lazyIntStream) Collect() []int {
	collectedData := make([]int, 0)
COLLECT_VALUE_LOOP:
	for _, value := range intStream.data {
		for _, function := range intStream.functions {
			castFilter, ok := function.(intFilter)
			if ok {
				if castFilter(value) {
					//Continue, since value fits the filter
					continue
				} else {
					//Value has to be sorted out, therefore skip to next value
					continue COLLECT_VALUE_LOOP
				}
			}

			castMapper, ok := function.(intMapper)
			if ok {
				value = castMapper(value)
				continue
			}

		}

		collectedData = append(collectedData, value)
	}

	return collectedData
}
