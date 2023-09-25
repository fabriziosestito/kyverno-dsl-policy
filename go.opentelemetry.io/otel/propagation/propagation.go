package propagation

import _ "unsafe"

type Embedme interface{}

type Baggage struct {}

func (b Baggage) Inject(ctx interface{}, carrier interface{}) {
 panic("stub")
}

func (b Baggage) Extract(parent interface{}, carrier interface{}) interface{} {
 panic("stub")
}

func (b Baggage) Fields() interface{} {
 panic("stub")
}

type TextMapCarrier interface {}

type MapCarrier map[string]string

type HeaderCarrier http.Header

type TextMapPropagator interface {}

type compositeTextMapPropagator []TextMapPropagator

func (c MapCarrier) Get(key interface{}) interface{} {
 panic("stub")
}

func (c MapCarrier) Set(key, value interface{}) {
 panic("stub")
}

func (c MapCarrier) Keys() interface{} {
 panic("stub")
}

func (hc HeaderCarrier) Get(key interface{}) interface{} {
 panic("stub")
}

func (hc HeaderCarrier) Set(key interface{}, value interface{}) {
 panic("stub")
}

func (hc HeaderCarrier) Keys() interface{} {
 panic("stub")
}

func (p compositeTextMapPropagator) Inject(ctx interface{}, carrier interface{}) {
 panic("stub")
}

func (p compositeTextMapPropagator) Extract(ctx interface{}, carrier interface{}) interface{} {
 panic("stub")
}

func (p compositeTextMapPropagator) Fields() interface{} {
 panic("stub")
}

func NewCompositeTextMapPropagator(p interface{}) interface{} {
 panic("stub")
}

type TraceContext struct {}

func (tc TraceContext) Inject(ctx interface{}, carrier interface{}) {
 panic("stub")
}

func (tc TraceContext) Extract(ctx interface{}, carrier interface{}) interface{} {
 panic("stub")
}

func (tc TraceContext) extract(carrier interface{}) interface{} {
 panic("stub")
}

func (tc TraceContext) Fields() interface{} {
 panic("stub")
}

