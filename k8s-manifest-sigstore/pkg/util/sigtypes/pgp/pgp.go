package pgp

import _ "unsafe"

type Embedme interface{}

func VerifyBlob(msgBytes, sigBytes interface{}, pubkeyPathString interface{}) (interface{}, interface{}, interface{}, interface{}) {
 panic("stub")
}

func LoadPublicKey(keyPath interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetFirstIdentity(signer interface{}) interface{} {
 panic("stub")
}

func getPublicKeyStream(keyRef interface{}) (interface{}, interface{}) {
 panic("stub")
}

