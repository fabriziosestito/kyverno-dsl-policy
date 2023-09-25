package cosign

import _ "unsafe"

type Embedme interface{}

type Cosign interface {}

type driver struct {}

func (d *driver) VerifyImageSignatures(ctx interface{}, signedImgRef interface{}, co interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func (d *driver) VerifyImageAttestations(ctx interface{}, signedImgRef interface{}, co interface{}) (checkedAttestations interface{}, bundleVerified interface{}, err interface{}) {
 panic("stub")
}

type cosignVerifier struct {}

func NewVerifier() interface{} {
 panic("stub")
}

func (v *cosignVerifier) VerifySignature(ctx interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func buildCosignOptions(ctx interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func loadCertPool(roots interface{}) (interface{}, interface{}) {
 panic("stub")
}

func loadCert(pem interface{}) (interface{}, interface{}) {
 panic("stub")
}

func loadCertChain(pem interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (v *cosignVerifier) FetchAttestations(ctx interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func matchType(sig interface{}, expectedType interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func decodeStatements(sigs interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func decodeStatement(sig interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func decodePayload(payloadBase64 interface{}) (interface{}, interface{}) {
 panic("stub")
}

func decodeCosignCustomProvenanceV01(statement interface{}) (interface{}, interface{}) {
 panic("stub")
}

func stringToJSONMap(i interface{}) (interface{}, interface{}) {
 panic("stub")
}

func decodePEM(raw interface{}, signatureAlgorithm interface{}) (interface{}, interface{}) {
 panic("stub")
}

func extractPayload(verified interface{}) (interface{}, interface{}) {
 panic("stub")
}

func extractDigest(imgRef interface{}, payload interface{}) (interface{}, interface{}) {
 panic("stub")
}

func matchSignatures(signatures interface{}, subject, issuer interface{}, extensions interface{}) interface{} {
 panic("stub")
}

func matchCertificateData(cert interface{}, subject, issuer interface{}, extensions interface{}) interface{} {
 panic("stub")
}

func matchExtensions(cert interface{}, issuer interface{}, extensions interface{}) interface{} {
 panic("stub")
}

func extractCertExtensionValue(key interface{}, ce interface{}) (interface{}, interface{}) {
 panic("stub")
}

func checkAnnotations(payload interface{}, annotations interface{}) interface{} {
 panic("stub")
}

func getRekorPubs(ctx interface{}, rekorPubKey interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getCTLogPubs(ctx interface{}, ctlogPubKey interface{}) (interface{}, interface{}) {
 panic("stub")
}

type mock struct {data interface{}}

func SetMock(image interface{}, data interface{}) interface{} {
 panic("stub")
}

func ClearMock() {
 panic("stub")
}

func (m *mock) VerifyImageSignatures(_ interface{}, signedImgRef interface{}, _ interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func (m *mock) VerifyImageAttestations(ctx interface{}, signedImgRef interface{}, co interface{}) (checkedAttestations interface{}, bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func (m *mock) getSignatures(signedImgRef interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func getSignature(sp interface{}) (interface{}, interface{}) {
 panic("stub")
}

