package x509

import _ "unsafe"

type Embedme interface{}

func VerifyBlob(msgBytes, sigBytes, certBytes interface{}, caCertPathString interface{}) (interface{}, interface{}, interface{}, interface{}) {
 panic("stub")
}

func LoadCertificate(certPath interface{}) (interface{}, interface{}) {
 panic("stub")
}

func LoadCertificateChain(certChainPath interface{}) (interface{}, interface{}) {
 panic("stub")
}

func PEMDecode(pemBytes interface{}, mode interface{}) interface{} {
 panic("stub")
}

func GetPublicKeyFromCertificate(certPemBytes interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetNameInfoFromX509Cert(cert interface{}) interface{} {
 panic("stub")
}

func isSelfSignedCert(cert interface{}) interface{} {
 panic("stub")
}

