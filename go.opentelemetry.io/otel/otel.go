package otel

import _ "unsafe"

type Embedme interface{}

type ErrorHandlerFunc func(interface{})

type ErrorHandler interface {}

func (f ErrorHandlerFunc) Handle(err interface{}) {
 panic("stub")
}

func GetErrorHandler() interface{} {
 panic("stub")
}

func SetErrorHandler(h interface{}) {
 panic("stub")
}

func Handle(err interface{}) {
 panic("stub")
}

func SetLogger(logger interface{}) {
 panic("stub")
}

func Meter(name interface{}, opts interface{}) interface{} {
 panic("stub")
}

func GetMeterProvider() interface{} {
 panic("stub")
}

func SetMeterProvider(mp interface{}) {
 panic("stub")
}

func GetTextMapPropagator() interface{} {
 panic("stub")
}

func SetTextMapPropagator(propagator interface{}) {
 panic("stub")
}

func Tracer(name interface{}, opts interface{}) interface{} {
 panic("stub")
}

func GetTracerProvider() interface{} {
 panic("stub")
}

func SetTracerProvider(tp interface{}) {
 panic("stub")
}

func Version() interface{} {
 panic("stub")
}

