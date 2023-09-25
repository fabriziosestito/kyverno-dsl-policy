package pkcs11key

import _ "unsafe"

type Embedme interface{}

type Key struct {}

type empty struct {}

func GetKeyWithURIConfig(config interface{}, askForPinIfNeeded interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) Certificate() (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) PublicKey(opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) VerifySignature(signature, message interface{}, opts interface{}) interface{} {
 panic("stub")
}

func (k *Key) Verifier() (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) Sign(ctx interface{}, rawPayload interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func (k *Key) SignMessage(message interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) SignerVerifier() (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) Close() {
 panic("stub")
}

type Pkcs11UriConfig struct {uriPathAttributes interface{}; uriQueryAttributes interface{}; ModulePath interface{}; SlotID interface{}; TokenLabel interface{}; KeyLabel interface{}; KeyID interface{}; Pin interface{}}

func percentEncode(input interface{}) interface{} {
 panic("stub")
}

func EncodeURIComponent(uriString interface{}, isForPath interface{}, usePercentEncoding interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewPkcs11UriConfig() interface{} {
 panic("stub")
}

func NewPkcs11UriConfigFromInput(modulePath interface{}, slotID interface{}, tokenLabel interface{}, keyLabel interface{}, keyID interface{}, pin interface{}) interface{} {
 panic("stub")
}

func (conf *Pkcs11UriConfig) Parse(uriString interface{}) interface{} {
 panic("stub")
}

func (conf *Pkcs11UriConfig) Construct() (interface{}, interface{}) {
 panic("stub")
}

