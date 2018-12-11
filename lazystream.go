package gostream

type lazyGenericStreamEntityStream struct {
	data      []GenericStreamEntity
	functions []interface{}
}

type GenericStreamEntityFilter func(value GenericStreamEntity) bool
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
