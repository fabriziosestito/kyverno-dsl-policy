package sigtypes

import _ "unsafe"

type Embedme interface{}

type SigType string

func GetSignatureTypeFromPublicKey(keyPathPtr interface{}) interface{} {
 panic("stub")
}

