package empty

import _ "unsafe"

type Embedme interface{}

type emptyImage struct {Embedme}

func Signatures() interface{} {
 panic("stub")
}

type signedImage struct {Embedme; digest interface{}; signature interface{}; attestations interface{}}

func (se *signedImage) Signatures() (interface{}, interface{}) {
 panic("stub")
}

func (se *signedImage) Attestations() (interface{}, interface{}) {
 panic("stub")
}

func (se *signedImage) Attachment(name interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (se *signedImage) Digest() (interface{}, interface{}) {
 panic("stub")
}

func SignedImage(ref interface{}) (interface{}, interface{}) {
 panic("stub")
}

