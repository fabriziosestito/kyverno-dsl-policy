package verify

import _ "unsafe"

type Embedme interface{}

type VerifyCommand struct {
	Embedme
	CheckClaims                  interface{}
	KeyRef                       interface{}
	CertRef                      interface{}
	CertGithubWorkflowTrigger    interface{}
	CertGithubWorkflowSha        interface{}
	CertGithubWorkflowName       interface{}
	CertGithubWorkflowRepository interface{}
	CertGithubWorkflowRef        interface{}
	CertChain                    interface{}
	CertOidcProvider             interface{}
	IgnoreSCT                    interface{}
	SCTRef                       interface{}
	Sk                           interface{}
	Slot                         interface{}
	Output                       interface{}
	RekorURL                     interface{}
	Attachment                   interface{}
	Annotations                  interface{}
	SignatureRef                 interface{}
	PayloadRef                   interface{}
	HashAlgorithm                interface{}
	LocalImage                   interface{}
	NameOptions                  interface{}
	Offline                      interface{}
	TSACertChainPath             interface{}
	IgnoreTlog                   interface{}
	MaxWorkers                   interface{}
}

func (c *VerifyCommand) Exec(ctx interface{}, images interface{}) interface{} {
	panic("stub")
}

func PrintVerificationHeader(ctx interface{}, imgRef interface{}, co interface{}, bundleVerified, fulcioVerified interface{}) {
	panic("stub")
}

func PrintVerification(ctx interface{}, verified interface{}, output interface{}) {
	panic("stub")
}

func loadCertFromFileOrURL(path interface{}) (interface{}, interface{}) {
	panic("stub")
}

func loadCertFromPEM(pems interface{}) (interface{}, interface{}) {
	panic("stub")
}

func loadCertChainFromFileOrURL(path interface{}) (interface{}, interface{}) {
	panic("stub")
}

func keylessVerification(keyRef interface{}, sk interface{}) interface{} {
	panic("stub")
}

type VerifyAttestationCommand struct {
	Embedme
	CheckClaims                  interface{}
	KeyRef                       interface{}
	CertRef                      interface{}
	CertGithubWorkflowTrigger    interface{}
	CertGithubWorkflowSha        interface{}
	CertGithubWorkflowName       interface{}
	CertGithubWorkflowRepository interface{}
	CertGithubWorkflowRef        interface{}
	CertChain                    interface{}
	IgnoreSCT                    interface{}
	SCTRef                       interface{}
	Sk                           interface{}
	Slot                         interface{}
	Output                       interface{}
	RekorURL                     interface{}
	PredicateType                interface{}
	Policies                     interface{}
	LocalImage                   interface{}
	NameOptions                  interface{}
	Offline                      interface{}
	TSACertChainPath             interface{}
	IgnoreTlog                   interface{}
	MaxWorkers                   interface{}
}

func (c *VerifyAttestationCommand) Exec(ctx interface{}, images interface{}) interface{} {
	panic("stub")
}

type VerifyBlobCmd struct {
	Embedme
	CertRef                      interface{}
	CertChain                    interface{}
	SigRef                       interface{}
	CertGithubWorkflowTrigger    interface{}
	CertGithubWorkflowSHA        interface{}
	CertGithubWorkflowName       interface{}
	CertGithubWorkflowRepository interface{}
	CertGithubWorkflowRef        interface{}
	IgnoreSCT                    interface{}
	SCTRef                       interface{}
	Offline                      interface{}
	IgnoreTlog                   interface{}
}

func isb64(data interface{}) interface{} {
	panic("stub")
}

func (c *VerifyBlobCmd) Exec(ctx interface{}, blobRef interface{}) interface{} {
	panic("stub")
}

func base64signature(sigRef, bundlePath interface{}) (interface{}, interface{}) {
	panic("stub")
}

func payloadBytes(blobRef interface{}) (interface{}, interface{}) {
	panic("stub")
}

type VerifyBlobAttestationCommand struct {
	Embedme
	CertRef                      interface{}
	CertChain                    interface{}
	CertGithubWorkflowTrigger    interface{}
	CertGithubWorkflowSHA        interface{}
	CertGithubWorkflowName       interface{}
	CertGithubWorkflowRepository interface{}
	CertGithubWorkflowRef        interface{}
	IgnoreSCT                    interface{}
	SCTRef                       interface{}
	Offline                      interface{}
	IgnoreTlog                   interface{}
	CheckClaims                  interface{}
	PredicateType                interface{}
	SignaturePath                interface{}
}

func (c *VerifyBlobAttestationCommand) Exec(ctx interface{}, artifactPath interface{}) interface{} {
	panic("stub")
}
