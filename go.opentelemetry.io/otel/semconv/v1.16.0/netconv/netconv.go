package netconv

import _ "unsafe"

type Embedme interface{}

func Transport(network interface{}) interface{} {
 panic("stub")
}

func Client(address interface{}, conn interface{}) interface{} {
 panic("stub")
}

func Server(address interface{}, ln interface{}) interface{} {
 panic("stub")
}

