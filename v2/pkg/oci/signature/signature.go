package signature

import _ "unsafe"

type Embedme interface{}

type sigLayer struct {Embedme; desc interface{}}

func New(l interface{}, desc interface{}) interface{} {
 panic("stub")
}

func (s *sigLayer) Annotations() (interface{}, interface{}) {
 panic("stub")
}

func (s *sigLayer) Payload() (interface{}, interface{}) {
 panic("stub")
}

func (s *sigLayer) Signature() (interface{}, interface{}) {
 panic("stub")
}

func (s *sigLayer) Base64Signature() (interface{}, interface{}) {
 panic("stub")
}

func (s *sigLayer) Cert() (interface{}, interface{}) {
 panic("stub")
}

func (s *sigLayer) Chain() (interface{}, interface{}) {
 panic("stub")
}

func (s *sigLayer) Bundle() (interface{}, interface{}) {
 panic("stub")
}

func (s *sigLayer) RFC3161Timestamp() (interface{}, interface{}) {
 panic("stub")
}

