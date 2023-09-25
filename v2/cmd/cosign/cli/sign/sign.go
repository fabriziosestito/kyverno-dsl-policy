package sign

import _ "unsafe"

type Embedme interface{}

type SignerVerifier struct {Cert interface{}; Chain interface{}; Embedme; close interface{}}

func ShouldUploadToTlog(ctx interface{}, ko interface{}, ref interface{}, tlogUpload interface{}) (interface{}, interface{}) {
 panic("stub")
}

func shouldUploadToTlog(ctx interface{}, ko interface{}, ref interface{}, tlogUpload interface{}) interface{} {
 panic("stub")
}

func GetAttachedImageRef(ref interface{}, attachment interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func ParseOCIReference(ctx interface{}, refStr interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func SignCmd(ro interface{}, ko interface{}, signOpts interface{}, imgs interface{}) interface{} {
 panic("stub")
}

func signDigest(ctx interface{}, digest interface{}, payload interface{}, ko interface{}, signOpts interface{}, annotations interface{}, dd interface{}, sv interface{}, se interface{}) interface{} {
 panic("stub")
}

func signerFromSecurityKey(ctx interface{}, keySlot interface{}) (interface{}, interface{}) {
 panic("stub")
}

func signerFromKeyRef(ctx interface{}, certPath, certChainPath, keyRef interface{}, passFunc interface{}) (interface{}, interface{}) {
 panic("stub")
}

func signerFromNewKey() (interface{}, interface{}) {
 panic("stub")
}

func keylessSigner(ctx interface{}, ko interface{}, sv interface{}) (interface{}, interface{}) {
 panic("stub")
}

func SignerFromKeyOpts(ctx interface{}, certPath interface{}, certChainPath interface{}, ko interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (c *SignerVerifier) Close() {
 panic("stub")
}

func (c *SignerVerifier) Bytes(ctx interface{}) (interface{}, interface{}) {
 panic("stub")
}

func SignBlobCmd(ro interface{}, ko interface{}, payloadPath interface{}, b64 interface{}, outputSignature interface{}, outputCertificate interface{}, tlogUpload interface{}) (interface{}, interface{}) {
 panic("stub")
}

func extractCertificate(ctx interface{}, sv interface{}) (interface{}, interface{}) {
 panic("stub")
}

