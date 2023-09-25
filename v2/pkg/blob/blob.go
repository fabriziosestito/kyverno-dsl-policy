package blob

import _ "unsafe"

type Embedme interface{}

type UnrecognizedSchemeError struct {Scheme interface{}}

func (e *UnrecognizedSchemeError) Error() interface{} {
 panic("stub")
}

func LoadFileOrURL(fileRef interface{}) (interface{}, interface{}) {
 panic("stub")
}

