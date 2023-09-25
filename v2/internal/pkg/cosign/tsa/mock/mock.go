package mock

import _ "unsafe"

type Embedme interface{}

type TSAClient struct {Embedme; Signer interface{}; CertChain interface{}; Time interface{}; Message interface{}}

type TSAClientOptions struct {Time interface{}; Message interface{}; Signer interface{}}

func NewTSAClient(o interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (c *TSAClient) GetTimestampResponse(tsq interface{}) (interface{}, interface{}) {
 panic("stub")
}

