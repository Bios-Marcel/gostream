package gostream

type eagerIntStream struct {
	data []int
}

//StreamIntsEager creates an eager IntStream that uses a copy of the passed
//array.
func StreamIntsEager(data []int) IntStream {
	defensiveCopy := make([]int, len(data))
	copy(defensiveCopy, data)
	return &eagerIntStream{
		data: defensiveCopy,
	}
}

func (intStream *eagerIntStream) Filter(filterFunction func(value int) bool) IntStream {
	newData := make([]int, 0)
	for _, element := range intStream.data {
		if filterFunction(element) {
			newData = append(newData, element)
		}
	}
	intStream.data = newData

	return intStream
}

func (intStream *eagerIntStream) Map(mapFunction func(value int) int) IntStream {
	for index, element := range intStream.data {
		intStream.data[index] = mapFunction(element)
	}

	return intStream
}

func (intStream *eagerIntStream) FindFirst() *int {
	if len(intStream.data) > 0 {
		return &intStream.data[0]
	}

	return nil
}

func (intStream *eagerIntStream) Collect() []int {
	return intStream.data
}
