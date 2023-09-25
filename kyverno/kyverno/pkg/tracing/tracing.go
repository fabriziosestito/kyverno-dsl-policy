package tracing

import _ "unsafe"

type Embedme interface{}

func StringValue(value interface{}) interface{} {
	panic("stub")
}

func StartChildSpan(ctx interface{}, tracerName interface{}, operationName interface{}, opts interface{}) (interface{}, interface{}) {
	panic("stub")
}

func ChildSpan(ctx interface{}, tracerName interface{}, operationName interface{}, doFn interface{}, opts ...interface{}) {
	panic("stub")
}

func ChildSpan1(ctx interface{}, tracerName interface{}, operationName interface{}, doFn interface{}, opts ...interface{}) interface{} {
	panic("stub")
}

func ChildSpan2(ctx interface{}, tracerName interface{}, operationName interface{}, doFn interface{}, opts ...interface{}) (interface{}, interface{}) {
	panic("stub")
}

func ChildSpan3(ctx interface{}, tracerName interface{}, operationName interface{}, doFn interface{}, opts ...interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func NewTraceConfig(log interface{}, tracerName, address, certs interface{}, kubeClient interface{}) (interface{}, interface{}) {
	panic("stub")
}

func SetSpanStatus(span interface{}, err interface{}) {
	panic("stub")
}

func SetStatus(ctx interface{}, err interface{}) {
	panic("stub")
}

func SetHttpStatus(ctx interface{}, err interface{}, code interface{}) {
	panic("stub")
}

func IsInSpan(ctx interface{}) interface{} {
	panic("stub")
}

func CurrentSpan(ctx interface{}) interface{} {
	panic("stub")
}

func SetAttributes(ctx interface{}, kv interface{}) {
	panic("stub")
}

func RequestFilterIsInSpan(request interface{}) interface{} {
	panic("stub")
}

func Transport(base interface{}, opts interface{}) interface{} {
	panic("stub")
}

func StartSpan(ctx interface{}, tracerName interface{}, operationName interface{}, opts interface{}) (interface{}, interface{}) {
	panic("stub")
}

func Span(ctx interface{}, tracerName interface{}, operationName interface{}, doFn interface{}, opts interface{}) {
	panic("stub")
}

func Span1(ctx interface{}, tracerName interface{}, operationName interface{}, doFn interface{}, opts interface{}) interface{} {
	panic("stub")
}

func Span2(ctx interface{}, tracerName interface{}, operationName interface{}, doFn interface{}, opts interface{}) (interface{}, interface{}) {
	panic("stub")
}

func Span3(ctx interface{}, tracerName interface{}, operationName interface{}, doFn interface{}, opts interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}
