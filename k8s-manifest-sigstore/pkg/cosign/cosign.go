package cosign

import _ "unsafe"

type Embedme interface{}

func SignImage(resBundleRef interface{}, keyPath, certPath interface{}, rekorURL interface{}, tlogUpload, force interface{}, pf interface{}, imageAnnotations interface{}, allowInsecure interface{}) interface{} {
 panic("stub")
}

func SignBlob(blobPath interface{}, keyPath, certPath interface{}, rekorURL interface{}, tlogUpload, force interface{}, pf interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetRekorServerURL() interface{} {
 panic("stub")
}

func GetTlogEntry(ctx interface{}, rekorClient interface{}, uuid interface{}) (interface{}, interface{}) {
 panic("stub")
}

func FindTLogEntriesByPayload(ctx interface{}, rekorClient interface{}, payload interface{}) (uuids interface{}, err interface{}) {
 panic("stub")
}

func ComputeLeafHash(e interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getUUID(entryUUID interface{}) (interface{}, interface{}) {
 panic("stub")
}

func verifyUUID(entryUUID interface{}, e interface{}) interface{} {
 panic("stub")
}

type cosignBundleSignature struct {Embedme; message interface{}; base64Signature interface{}; cert interface{}; base64Bundle interface{}}

func createCosignCheckOpts(pubkeyPath, certRef, certChain, rekorURL, oidcIssuer interface{}, rootCerts interface{}, allowInsecure interface{}) (interface{}, interface{}) {
 panic("stub")
}

func VerifyImage(resBundleRef, pubkeyPath, certRef, certChain, rekorURL, oidcIssuer interface{}, rootCerts interface{}, allowInsecure interface{}) (interface{}, interface{}, interface{}, interface{}) {
 panic("stub")
}

func VerifyBlob(msgBytes, sigBytes, certBytes, bundleBytes interface{}, pubkeyPath interface{}, certRef, certChain, rekorURL, oidcIssuer interface{}, rootCerts interface{}) (interface{}, interface{}, interface{}, interface{}) {
 panic("stub")
}

func loadCertificate(pemBytes interface{}) (interface{}, interface{}) {
 panic("stub")
}

func verifyBundle(rawMsg, b64Sig, rawCert, b64Bundle interface{}, co interface{}) (interface{}, interface{}, interface{}, interface{}) {
 panic("stub")
}

func (s *cosignBundleSignature) Annotations() (interface{}, interface{}) {
 panic("stub")
}

func (s *cosignBundleSignature) Payload() (interface{}, interface{}) {
 panic("stub")
}

func (s *cosignBundleSignature) Signature() (interface{}, interface{}) {
 panic("stub")
}

func (s *cosignBundleSignature) Base64Signature() (interface{}, interface{}) {
 panic("stub")
}

func (s *cosignBundleSignature) Cert() (interface{}, interface{}) {
 panic("stub")
}

func (s *cosignBundleSignature) Chain() (interface{}, interface{}) {
 panic("stub")
}

func (s *cosignBundleSignature) Bundle() (interface{}, interface{}) {
 panic("stub")
}

func (s *cosignBundleSignature) RFC3161Timestamp() (interface{}, interface{}) {
 panic("stub")
}

