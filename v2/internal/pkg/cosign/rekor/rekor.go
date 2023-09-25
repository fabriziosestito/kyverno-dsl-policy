package rekor

import _ "unsafe"

type Embedme interface{}

type tlogUploadFn func(interface{}, interface{}) (interface{}, interface{})

type signerWrapper struct {inner interface{}; rClient interface{}}

func uploadToTlog(rekorBytes interface{}, rClient interface{}, upload interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (rs *signerWrapper) Sign(ctx interface{}, payload interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func NewSigner(inner interface{}, rClient interface{}) interface{} {
 panic("stub")
}

