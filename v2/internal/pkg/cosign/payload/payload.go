package payload

import _ "unsafe"

type Embedme interface{}

type payloadAttestor struct {signer interface{}; payloadType interface{}}

func (pa *payloadAttestor) DSSEAttest(ctx interface{}, payload interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func NewDSSEAttestor(payloadType interface{}, s interface{}, signAndPublicKeyOptions interface{}) interface{} {
 panic("stub")
}

type payloadSigner struct {payloadSigner interface{}; payloadSignerOpts interface{}; publicKeyProviderOpts interface{}}

func (ps *payloadSigner) Sign(ctx interface{}, payload interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func (ps *payloadSigner) publicKey(ctx interface{}) (pk interface{}, err interface{}) {
 panic("stub")
}

func (ps *payloadSigner) signPayload(ctx interface{}, payloadBytes interface{}) (sig interface{}, err interface{}) {
 panic("stub")
}

func newSigner(s interface{}, signAndPublicKeyOptions interface{}) interface{} {
 panic("stub")
}

func NewSigner(s interface{}, signAndPublicKeyOptions interface{}) interface{} {
 panic("stub")
}

