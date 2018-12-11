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
