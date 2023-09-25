package pivkey

import _ "unsafe"

type Embedme interface{}

type empty struct {}

type Key struct {}

func GetKey() (interface{}, interface{}) {
 panic("stub")
}

func GetKeyWithSlot(slot interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) Close() {
 panic("stub")
}

func (k *Key) Authenticate(pin interface{}) {
 panic("stub")
}

func (k *Key) SetSlot(slot interface{}) {
 panic("stub")
}

func (k *Key) Attest() (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) GetAttestationCertificate() (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) SetManagementKey(old, new interface{}) interface{} {
 panic("stub")
}

func (k *Key) SetPIN(old, new interface{}) interface{} {
 panic("stub")
}

func (k *Key) SetPUK(old, new interface{}) interface{} {
 panic("stub")
}

func (k *Key) Reset() interface{} {
 panic("stub")
}

func (k *Key) Unblock(puk, newPIN interface{}) interface{} {
 panic("stub")
}

func (k *Key) GenerateKey(mgmtKey interface{}, slot interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) Verifier() (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) Certificate() (interface{}, interface{}) {
 panic("stub")
}

func (k *Key) SignerVerifier() (interface{}, interface{}) {
 panic("stub")
}

