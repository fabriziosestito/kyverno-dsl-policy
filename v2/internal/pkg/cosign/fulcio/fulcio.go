package fulcio

import _ "unsafe"

type Embedme interface{}

type signerWrapper struct {inner interface{}; cert; chain interface{}}

func (fs *signerWrapper) Sign(ctx interface{}, payload interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func NewSigner(inner interface{}, cert, chain interface{}) interface{} {
 panic("stub")
}

