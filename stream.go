package gostream

import "github.com/cheekybits/genny/generic"

//GenericStreamEntity is the generic type that will be replaced on code
//generation for your needed types.
type GenericStreamEntity generic.Type

//GenericStreamEntityStream is a stream over an array of
//GenericStreamEntityStreams. Depending on its implemenation its behaviour may
//vary. The default implementations are an eager stream and a lazy stream.
type GenericStreamEntityStream interface {
	//Filter only keeps the values that meet the filter condition.
	Filter(func(value GenericStreamEntity) bool) GenericStreamEntityStream
	//Map changes the values according to the mapping function.
	Map(func(value GenericStreamEntity) GenericStreamEntity) GenericStreamEntityStream
	//Reduces combines all values using the given reduce function.
	Reduce(func(one, two GenericStreamEntity) GenericStreamEntity) *GenericStreamEntity
	//FindFirst returns the first element found.
	FindFirst() *GenericStreamEntity
	//Collect returns a copy of the internally used int array.
	Collect() []GenericStreamEntity
}

//StreamGenericStreamEntity creates a StreamGenericStreamEntity using the lazy
//implementation of StreamGenericStreamEntityStream.
func StreamGenericStreamEntity(data []GenericStreamEntity) GenericStreamEntityStream {
	return StreamGenericStreamEntityLazy(data)
}

type lazyGenericStreamEntityStream struct {
	data      []GenericStreamEntity
	functions []interface{}
}

//GenericStreamEntityFilter is the type of function that filters values.
type GenericStreamEntityFilter func(value GenericStreamEntity) bool

//GenericStreamEntityMapper is the type of function that maps values.
type GenericStreamEntityMapper func(value GenericStreamEntity) GenericStreamEntity

//StreamGenericStreamEntityLazy creates a lazy StreamGenericStreamEntity that
//uses a copy of the passed array.
func StreamGenericStreamEntityLazy(data []GenericStreamEntity) GenericStreamEntityStream {
	defensiveCopy := make([]GenericStreamEntity, len(data))
	copy(defensiveCopy, data)
	return &lazyGenericStreamEntityStream{
		data: defensiveCopy,
	}
}

func (stream *lazyGenericStreamEntityStream) Filter(filterFunction func(value GenericStreamEntity) bool) GenericStreamEntityStream {
	stream.functions = append(stream.functions, (GenericStreamEntityFilter)(filterFunction))
	return stream
}

func (stream *lazyGenericStreamEntityStream) Map(mapFunction func(value GenericStreamEntity) GenericStreamEntity) GenericStreamEntityStream {
	stream.functions = append(stream.functions, (GenericStreamEntityMapper)(mapFunction))
	return stream
}

func (stream *lazyGenericStreamEntityStream) FindFirst() *GenericStreamEntity {
FINDFIRST_VALUE_LOOP:
	for _, value := range stream.data {
		for _, function := range stream.functions {
			castFilter, ok := function.(GenericStreamEntityFilter)
			if ok {
				if castFilter(value) {
					//Continue, since value fits the filter
					continue
				} else {
					//Value has to be sorted out, therefore skip to next value
					continue FINDFIRST_VALUE_LOOP
				}
			}

			castMapper, ok := function.(GenericStreamEntityMapper)
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

func (stream *lazyGenericStreamEntityStream) Reduce(reduceFunction func(one, two GenericStreamEntity) GenericStreamEntity) *GenericStreamEntity {
	//This implementation is just me being lazy, instead of a proper
	//optimized solution, I'll simply call collect and reduce the result.
	return reduceGenericStreamEntity(reduceFunction, stream.Collect())
}

func (stream *lazyGenericStreamEntityStream) Collect() []GenericStreamEntity {
	collectedData := make([]GenericStreamEntity, 0)
COLLECT_VALUE_LOOP:
	for _, value := range stream.data {
		for _, function := range stream.functions {
			castFilter, ok := function.(GenericStreamEntityFilter)
			if ok {
				if castFilter(value) {
					//Continue, since value fits the filter
					continue
				} else {
					//Value has to be sorted out, therefore skip to next value
					continue COLLECT_VALUE_LOOP
				}
			}

			castMapper, ok := function.(GenericStreamEntityMapper)
			if ok {
				value = castMapper(value)
				continue
			}

		}

		collectedData = append(collectedData, value)
	}

	return collectedData
}

type eagerGenericStreamEntityStream struct {
	data []GenericStreamEntity
}

//StreamGenericStreamEntityEager creates an eager GenericStreamEntityStream
//that uses a copy of the passed array.
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

func reduceGenericStreamEntity(reduceFunction func(valueOne, valueTwo GenericStreamEntity) GenericStreamEntity, data []GenericStreamEntity) *GenericStreamEntity {
	lengthOfData := len(data)

	if lengthOfData == 0 {
		return nil
	}

	if lengthOfData == 1 {
		return &data[0]
	}

	var value GenericStreamEntity
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
