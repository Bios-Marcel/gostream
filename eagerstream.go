package gostream

type eagerGenericStreamEntityStream struct {
	data []GenericStreamEntity
}

//StreamIntsEager creates an eager GenericStreamEntityStream that uses a copy of the passed
//array.
func StreamGenericStreamEntityEager(data []GenericStreamEntity) GenericStreamEntityStream {
	defensiveCopy := make([]GenericStreamEntity, len(data))
	copy(defensiveCopy, data)
	return &eagerGenericStreamEntityStream{
		data: defensiveCopy,
	}
}

func (stream *eagerGenericStreamEntityStream) Filter(filterFunction func(value GenericStreamEntity) bool) GenericStreamEntityStream {
	newData := make([]GenericStreamEntity, 0)
	for _, element := range stream.data {
		if filterFunction(element) {
			newData = append(newData, element)
		}
	}
	stream.data = newData

	return stream
}

func (stream *eagerGenericStreamEntityStream) Map(mapFunction func(value GenericStreamEntity) GenericStreamEntity) GenericStreamEntityStream {
	for index, element := range stream.data {
		stream.data[index] = mapFunction(element)
	}

	return stream
}

func (stream *eagerGenericStreamEntityStream) FindFirst() *GenericStreamEntity {
	if len(stream.data) > 0 {
		return &stream.data[0]
	}

	return nil
}

func (stream *eagerGenericStreamEntityStream) Collect() []GenericStreamEntity {
	return stream.data
}

func (stream *eagerGenericStreamEntityStream) Reduce(reduceFunction func(valueOne, valueTwo GenericStreamEntity) GenericStreamEntity) *GenericStreamEntity {
	return reduceGenericStreamEntity(reduceFunction, stream.data)
}
