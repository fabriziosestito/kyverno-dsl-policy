package cosign

import _ "unsafe"

type Embedme interface{}

type CertExtensions struct {Cert interface{}}

func (ce *CertExtensions) certExtensions() interface{} {
 panic("stub")
}

func (ce *CertExtensions) GetIssuer() interface{} {
 panic("stub")
}

func (ce *CertExtensions) GetCertExtensionGithubWorkflowTrigger() interface{} {
 panic("stub")
}

func (ce *CertExtensions) GetExtensionGithubWorkflowSha() interface{} {
 panic("stub")
}

func (ce *CertExtensions) GetCertExtensionGithubWorkflowName() interface{} {
 panic("stub")
}

func (ce *CertExtensions) GetCertExtensionGithubWorkflowRepository() interface{} {
 panic("stub")
}

func (ce *CertExtensions) GetCertExtensionGithubWorkflowRef() interface{} {
 panic("stub")
}

func GetPassFromTerm(confirm interface{}) (interface{}, interface{}) {
 panic("stub")
}

func IsTerminal() interface{} {
 panic("stub")
}

func GetCTLogPubs(ctx interface{}) (interface{}, interface{}) {
 panic("stub")
}

type VerificationFailure struct {err interface{}}

type ErrNoMatchingSignatures struct {err interface{}}

type ErrImageTagNotFound struct {err interface{}}

type ErrNoMatchingAttestations struct {err interface{}}

type ErrNoCertificateFoundOnSignature struct {err interface{}}

type ErrNoSignaturesFound struct {err interface{}}

type VerificationError struct {message interface{}}

func (e *VerificationFailure) Error() interface{} {
 panic("stub")
}

func (e *VerificationFailure) Unwrap() interface{} {
 panic("stub")
}

func (e *ErrNoMatchingSignatures) Error() interface{} {
 panic("stub")
}

func (e *ErrNoMatchingSignatures) Unwrap() interface{} {
 panic("stub")
}

func (e *ErrImageTagNotFound) Error() interface{} {
 panic("stub")
}

func (e *ErrImageTagNotFound) Unwrap() interface{} {
 panic("stub")
}

func (e *ErrNoSignaturesFound) Error() interface{} {
 panic("stub")
}

func (e *ErrNoSignaturesFound) Unwrap() interface{} {
 panic("stub")
}

func (e *ErrNoMatchingAttestations) Error() interface{} {
 panic("stub")
}

func (e *ErrNoMatchingAttestations) Unwrap() interface{} {
 panic("stub")
}

func (e *ErrNoCertificateFoundOnSignature) Error() interface{} {
 panic("stub")
}

func (e *ErrNoCertificateFoundOnSignature) Unwrap() interface{} {
 panic("stub")
}

func NewVerificationError(msg interface{}, args interface{}) interface{} {
 panic("stub")
}

func (e *VerificationError) Error() interface{} {
 panic("stub")
}

type SignedPayload struct {Base64Signature interface{}; Payload interface{}; Cert interface{}; Chain interface{}; Bundle interface{}; RFC3161Timestamp interface{}}

type LocalSignedPayload struct {Base64Signature interface{}; Cert interface{}; Bundle interface{}}

type Signatures struct {KeyID interface{}; Sig interface{}}

type AttestationPayload struct {PayloadType interface{}; PayLoad interface{}; Signatures interface{}}

func FetchSignaturesForReference(_ interface{}, ref interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func FetchAttestationsForReference(_ interface{}, ref interface{}, predicateType interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func FetchAttestations(se interface{}, predicateType interface{}) (interface{}, interface{}) {
 panic("stub")
}

func FetchLocalSignedPayloadFromPath(path interface{}) (interface{}, interface{}) {
 panic("stub")
}

type PassFunc func(interface{}) (interface{}, interface{})

type KeysBytes struct {PrivateBytes interface{}; PublicBytes interface{}; password interface{}}

type Keys struct {private interface{}; public interface{}}

func (k *KeysBytes) Password() interface{} {
 panic("stub")
}

func GeneratePrivateKey() (interface{}, interface{}) {
 panic("stub")
}

func ImportKeyPair(keyPath interface{}, pf interface{}) (interface{}, interface{}) {
 panic("stub")
}

func marshalKeyPair(ptype interface{}, keypair interface{}, pf interface{}) (key interface{}, err interface{}) {
 panic("stub")
}

func GenerateKeyPair(pf interface{}) (interface{}, interface{}) {
 panic("stub")
}

func PemToECDSAKey(pemBytes interface{}) (interface{}, interface{}) {
 panic("stub")
}

func LoadPrivateKey(key interface{}, pass interface{}) (interface{}, interface{}) {
 panic("stub")
}

func ObsoletePayload(ctx interface{}, digestedImage interface{}) (interface{}, interface{}) {
 panic("stub")
}

type key struct {}

func Set(ctx interface{}, rekorClient interface{}) interface{} {
 panic("stub")
}

func Get(ctx interface{}) interface{} {
 panic("stub")
}

type TrustedTransparencyLogPubKeys struct {Keys interface{}}

type TransparencyLogPubKey struct {PubKey interface{}; Status interface{}}

func GetTransparencyLogID(pub interface{}) (interface{}, interface{}) {
 panic("stub")
}

func dsseEntry(ctx interface{}, signature, pubKey interface{}) (interface{}, interface{}) {
 panic("stub")
}

func intotoEntry(ctx interface{}, signature, pubKey interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetRekorPubs(ctx interface{}) (interface{}, interface{}) {
 panic("stub")
}

func rekorPubsFromClient(rekorClient interface{}) (interface{}, interface{}) {
 panic("stub")
}

func TLogUpload(ctx interface{}, rekorClient interface{}, signature interface{}, sha256CheckSum interface{}, pemBytes interface{}) (interface{}, interface{}) {
 panic("stub")
}

func TLogUploadDSSEEnvelope(ctx interface{}, rekorClient interface{}, signature, pemBytes interface{}) (interface{}, interface{}) {
 panic("stub")
}

func TLogUploadInTotoAttestation(ctx interface{}, rekorClient interface{}, signature, pemBytes interface{}) (interface{}, interface{}) {
 panic("stub")
}

func doUpload(ctx interface{}, rekorClient interface{}, pe interface{}) (interface{}, interface{}) {
 panic("stub")
}

func rekorEntry(sha256CheckSum interface{}, signature, pubKey interface{}) interface{} {
 panic("stub")
}

func ComputeLeafHash(e interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getUUID(entryUUID interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getTreeUUID(entryUUID interface{}) (interface{}, interface{}) {
 panic("stub")
}

func isExpectedResponseUUID(requestEntryUUID interface{}, responseEntryUUID interface{}, treeid interface{}) interface{} {
 panic("stub")
}

func verifyUUID(entryUUID interface{}, e interface{}) interface{} {
 panic("stub")
}

func GetTlogEntry(ctx interface{}, rekorClient interface{}, entryUUID interface{}) (interface{}, interface{}) {
 panic("stub")
}

func proposedEntries(b64Sig interface{}, payload, pubKey interface{}) (interface{}, interface{}) {
 panic("stub")
}

func FindTlogEntry(ctx interface{}, rekorClient interface{}, b64Sig interface{}, payload, pubKey interface{}) (interface{}, interface{}) {
 panic("stub")
}

func VerifyTLogEntryOffline(ctx interface{}, e interface{}, rekorPubKeys interface{}) interface{} {
 panic("stub")
}

func NewTrustedTransparencyLogPubKeys() interface{} {
 panic("stub")
}

func (t *TrustedTransparencyLogPubKeys) AddTransparencyLogPubKey(pemBytes interface{}, status interface{}) interface{} {
 panic("stub")
}

func SimpleClaimVerifier(sig interface{}, imageDigest interface{}, annotations interface{}) interface{} {
 panic("stub")
}

func IntotoSubjectClaimVerifier(sig interface{}, imageDigest interface{}, _ interface{}) interface{} {
 panic("stub")
}

type CheckOpts struct {RegistryClientOpts interface{}; Annotations interface{}; ClaimVerifier interface{}; RekorClient interface{}; RekorPubKeys interface{}; SigVerifier interface{}; PKOpts interface{}; RootCerts interface{}; IntermediateCerts interface{}; CertGithubWorkflowTrigger interface{}; CertGithubWorkflowSha interface{}; CertGithubWorkflowName interface{}; CertGithubWorkflowRepository interface{}; CertGithubWorkflowRef interface{}; IgnoreSCT interface{}; SCT interface{}; CTLogPubKeys interface{}; SignatureRef interface{}; PayloadRef interface{}; Identities interface{}; Offline interface{}; TSACertificate interface{}; TSARootCertificates interface{}; TSAIntermediateCertificates interface{}; IgnoreTlog interface{}; MaxWorkers interface{}}

type fakeOCISignatures struct {Embedme; signatures interface{}}

type payloader interface {}

type Identity struct {Issuer interface{}; Subject interface{}; IssuerRegExp interface{}; SubjectRegExp interface{}}

type signatureVerificationFn func(ctx interface{}, verifier interface{}, sig interface{}) interface{}

func verifyOCIAttestation(ctx interface{}, verifier interface{}, att interface{}) interface{} {
 panic("stub")
}

func verifyOCISignature(ctx interface{}, verifier interface{}, sig interface{}) interface{} {
 panic("stub")
}

func ValidateAndUnpackCert(cert interface{}, co interface{}) (interface{}, interface{}) {
 panic("stub")
}

func CheckCertificatePolicy(cert interface{}, co interface{}) interface{} {
 panic("stub")
}

func validateCertExtensions(ce interface{}, co interface{}) interface{} {
 panic("stub")
}

func ValidateAndUnpackCertWithChain(cert interface{}, chain interface{}, co interface{}) (interface{}, interface{}) {
 panic("stub")
}

func tlogValidateEntry(ctx interface{}, client interface{}, rekorPubKeys interface{}, sig interface{}, pem interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (fos *fakeOCISignatures) Get() (interface{}, interface{}) {
 panic("stub")
}

func VerifyImageSignatures(ctx interface{}, signedImgRef interface{}, co interface{}) (checkedSignatures interface{}, bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func VerifyLocalImageSignatures(ctx interface{}, path interface{}, co interface{}) (checkedSignatures interface{}, bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func verifySignatures(ctx interface{}, sigs interface{}, h interface{}, co interface{}) (checkedSignatures interface{}, bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func verifyInternal(ctx interface{}, sig interface{}, h interface{}, verifyFn interface{}, co interface{}) (bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func keyBytes(sig interface{}, co interface{}) (interface{}, interface{}) {
 panic("stub")
}

func VerifyBlobSignature(ctx interface{}, sig interface{}, co interface{}) (bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func VerifyImageSignature(ctx interface{}, sig interface{}, h interface{}, co interface{}) (bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func loadSignatureFromFile(ctx interface{}, sigRef interface{}, signedImgRef interface{}, co interface{}) (interface{}, interface{}) {
 panic("stub")
}

func VerifyImageAttestations(ctx interface{}, signedImgRef interface{}, co interface{}) (checkedAttestations interface{}, bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func VerifyLocalImageAttestations(ctx interface{}, path interface{}, co interface{}) (checkedAttestations interface{}, bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func VerifyBlobAttestation(ctx interface{}, att interface{}, h interface{}, co interface{}) (interface{}, interface{}) {
 panic("stub")
}

func VerifyImageAttestation(ctx interface{}, atts interface{}, h interface{}, co interface{}) (checkedAttestations interface{}, bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func CheckExpiry(cert interface{}, it interface{}) interface{} {
 panic("stub")
}

func getBundleIntegratedTime(sig interface{}) (interface{}, interface{}) {
 panic("stub")
}

func VerifyBundle(sig interface{}, co interface{}) (interface{}, interface{}) {
 panic("stub")
}

func VerifyRFC3161Timestamp(sig interface{}, co interface{}) (interface{}, interface{}) {
 panic("stub")
}

func compareSigs(bundleBody interface{}, sig interface{}) interface{} {
 panic("stub")
}

func comparePublicKey(bundleBody interface{}, sig interface{}, co interface{}) interface{} {
 panic("stub")
}

func extractEntryImpl(bundleBody interface{}) (interface{}, interface{}) {
 panic("stub")
}

func bundleHash(bundleBody, _ interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func bundleSig(bundleBody interface{}) (interface{}, interface{}) {
 panic("stub")
}

func bundleKey(bundleBody interface{}) (interface{}, interface{}) {
 panic("stub")
}

func VerifySET(bundlePayload interface{}, signature interface{}, pub interface{}) interface{} {
 panic("stub")
}

func TrustedCert(cert interface{}, roots interface{}, intermediates interface{}) (interface{}, interface{}) {
 panic("stub")
}

func correctAnnotations(wanted, have interface{}) interface{} {
 panic("stub")
}

func verifyImageSignaturesExperimentalOCI(ctx interface{}, signedImgRef interface{}, co interface{}) (checkedSignatures interface{}, bundleVerified interface{}, err interface{}) {
 panic("stub")
}

func ContainsSCT(cert interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getCTPublicKey(sct interface{}, pubKeys interface{}) (interface{}, interface{}) {
 panic("stub")
}

func VerifySCT(_ interface{}, certPEM, chainPEM, rawSCT interface{}, pubKeys interface{}) interface{} {
 panic("stub")
}

func VerifyEmbeddedSCT(ctx interface{}, chain interface{}, pubKeys interface{}) interface{} {
 panic("stub")
}

