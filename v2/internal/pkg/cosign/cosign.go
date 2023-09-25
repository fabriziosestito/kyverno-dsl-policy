package cosign

import _ "unsafe"

type Embedme interface{}

type HashReader struct {r interface{}; h interface{}}

func FileExists(filename interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewHashReader(r interface{}, h interface{}) interface{} {
 panic("stub")
}

func (h *HashReader) Read(p interface{}) (n interface{}, err interface{}) {
 panic("stub")
}

func (h *HashReader) Sum(p interface{}) interface{} {
 panic("stub")
}

func (h *HashReader) Reset() {
 panic("stub")
}

func (h *HashReader) Size() interface{} {
 panic("stub")
}

func (h *HashReader) BlockSize() interface{} {
 panic("stub")
}

func (h *HashReader) Write(p interface{}) (interface{}, interface{}) {
 panic("stub")
}

type DSSEAttestor interface {}

type Signer interface {}

