package gostream

type eagerIntStream struct {
	data *[]int
}

func (intStream *eagerIntStream) getData() *[]int {
	return intStream.data
}

func (intStream *eagerIntStream) setData(newData *[]int) {
	intStream.data = newData
}

//StreamIntsEager creates an eager IntStream.
func StreamIntsEager(data []int) IntStream {
	return &eagerIntStream{
		data: &data,
	}
}

func (intStream *eagerIntStream) Filter(filterFunction func(value int) bool) IntStream {
	newData := make([]int, 0)
	for _, element := range *intStream.getData() {
		if filterFunction(element) {
			newData = append(newData, element)
		}
	}
	intStream.setData(&newData)

	return intStream
}

func (intStream *eagerIntStream) Map(mapFunction func(value int) int) IntStream {
	data := *intStream.getData()
	for index, element := range data {
		data[index] = mapFunction(element)
	}

	return intStream
}

func (intStream *eagerIntStream) FindFirst() *int {
	data := *intStream.getData()
	if len(data) > 0 {
		return &data[0]
	}

	return nil
}

func (intStream *eagerIntStream) Collect() []int {
	return *intStream.getData()
}
