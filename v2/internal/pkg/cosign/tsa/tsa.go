package tsa

import _ "unsafe"

type Embedme interface{}

type signerWrapper struct {inner interface{}; tsaClient interface{}}

func GetTimestampedSignature(sigBytes interface{}, tsaClient interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (rs *signerWrapper) Sign(ctx interface{}, payload interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func createTimestampAuthorityRequest(artifactBytes interface{}, hash interface{}, policyStr interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewSigner(inner interface{}, tsaClient interface{}) interface{} {
 panic("stub")
}

func SplitPEMCertificateChain(pem interface{}) (leaves, intermediates, roots interface{}, err interface{}) {
 panic("stub")
}

