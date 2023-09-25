package signature

import _ "unsafe"

type Embedme interface{}

type AnnotationsMap struct {Annotations interface{}}

func (a *AnnotationsMap) Set(s interface{}) interface{} {
 panic("stub")
}

func (a *AnnotationsMap) String() interface{} {
 panic("stub")
}

func LoadPublicKey(ctx interface{}, keyRef interface{}) (verifier interface{}, err interface{}) {
 panic("stub")
}

func VerifierForKeyRef(ctx interface{}, keyRef interface{}, hashAlgorithm interface{}) (verifier interface{}, err interface{}) {
 panic("stub")
}

func loadKey(keyPath interface{}, pf interface{}) (interface{}, interface{}) {
 panic("stub")
}

func LoadPublicKeyRaw(raw interface{}, hashAlgorithm interface{}) (interface{}, interface{}) {
 panic("stub")
}

func SignerFromKeyRef(ctx interface{}, keyRef interface{}, pf interface{}) (interface{}, interface{}) {
 panic("stub")
}

func SignerVerifierFromKeyRef(ctx interface{}, keyRef interface{}, pf interface{}) (interface{}, interface{}) {
 panic("stub")
}

func PublicKeyFromKeyRef(ctx interface{}, keyRef interface{}) (interface{}, interface{}) {
 panic("stub")
}

func PublicKeyFromKeyRefWithHashAlgo(ctx interface{}, keyRef interface{}, hashAlgorithm interface{}) (interface{}, interface{}) {
 panic("stub")
}

func PublicKeyPem(key interface{}, pkOpts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func CertSubject(c interface{}) interface{} {
 panic("stub")
}

