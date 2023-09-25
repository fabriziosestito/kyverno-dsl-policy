package baggage

import _ "unsafe"

type Embedme interface{}

type Item struct {Value interface{}; Properties interface{}}

type Property struct {Key; Value interface{}; HasValue interface{}}

type List map[string]Item

type baggageContextKeyType int

type GetHookFunc func(interface{}, interface{}) interface{}

type baggageState struct {list interface{}; setHook interface{}; getHook interface{}}

type SetHookFunc func(interface{}, interface{}) interface{}

func ContextWithSetHook(parent interface{}, hook interface{}) interface{} {
 panic("stub")
}

func ContextWithGetHook(parent interface{}, hook interface{}) interface{} {
 panic("stub")
}

func ContextWithList(parent interface{}, list interface{}) interface{} {
 panic("stub")
}

func ListFromContext(ctx interface{}) interface{} {
 panic("stub")
}

