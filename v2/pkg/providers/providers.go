package providers

import _ "unsafe"

type Embedme interface{}

type providerEntry struct {name interface{}; p interface{}}

type Interface interface {}

func Register(name interface{}, p interface{}) {
 panic("stub")
}

func Enabled(ctx interface{}) interface{} {
 panic("stub")
}

func Provide(ctx interface{}, audience interface{}) (interface{}, interface{}) {
 panic("stub")
}

func ProvideFrom(_ interface{}, provider interface{}) (interface{}, interface{}) {
 panic("stub")
}

