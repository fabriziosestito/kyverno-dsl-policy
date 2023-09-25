package github

import _ "unsafe"

type Embedme interface{}

type Gh struct {}

func New() interface{} {
 panic("stub")
}

func (g *Gh) PutSecret(ctx interface{}, ref interface{}, pf interface{}) interface{} {
 panic("stub")
}

func (g *Gh) GetSecret(ctx interface{}, ref interface{}, key interface{}) (interface{}, interface{}) {
 panic("stub")
}

func encryptSecretWithPublicKey(publicKey interface{}, secretName interface{}, secretValue interface{}) (interface{}, interface{}) {
 panic("stub")
}

