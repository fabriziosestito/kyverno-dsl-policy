package ephemeral

import _ "unsafe"

type Embedme interface{}

type ephemeralSigner struct {signer interface{}}

func (ks ephemeralSigner) Sign(_ interface{}, payload interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func NewSigner() (interface{}, interface{}) {
 panic("stub")
}

