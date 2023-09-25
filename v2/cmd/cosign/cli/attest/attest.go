package attest

import _ "unsafe"

type Embedme interface{}

type AttestCommand struct {Embedme; Embedme; CertPath interface{}; CertChainPath interface{}; NoUpload interface{}; PredicatePath interface{}; PredicateType interface{}; Replace interface{}; Timeout interface{}; TlogUpload interface{}; TSAServerURL interface{}}

type tlogUploadFn func(interface{}, interface{}) (interface{}, interface{})

func uploadToTlog(ctx interface{}, sv interface{}, rekorURL interface{}, upload interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (c *AttestCommand) Exec(ctx interface{}, imageRef interface{}) interface{} {
 panic("stub")
}

type AttestBlobCommand struct {Embedme; CertPath interface{}; CertChainPath interface{}; ArtifactHash interface{}; PredicatePath interface{}; PredicateType interface{}; TlogUpload interface{}; Timeout interface{}; OutputSignature interface{}; OutputAttestation interface{}; OutputCertificate interface{}}

func (c *AttestBlobCommand) Exec(ctx interface{}, artifactPath interface{}) interface{} {
 panic("stub")
}

func predicateReader(predicatePath interface{}) (interface{}, interface{}) {
 panic("stub")
}

