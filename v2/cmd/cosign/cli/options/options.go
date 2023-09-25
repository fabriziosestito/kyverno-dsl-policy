package options

import _ "unsafe"

type Embedme interface{}

type AnnotationOptions struct{ Annotations interface{} }

func (o *AnnotationOptions) AnnotationsMap() (interface{}, interface{}) {
	panic("stub")
}

func (o *AnnotationOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type AttachAttestationOptions struct {
	Attestations interface{}
	Registry     interface{}
}

type AttachSignatureOptions struct {
	Signature      interface{}
	Payload        interface{}
	Cert           interface{}
	CertChain      interface{}
	TimeStampedSig interface{}
	Registry       interface{}
}

type AttachSBOMOptions struct {
	SBOM                 interface{}
	SBOMType             interface{}
	SBOMInputFormat      interface{}
	Registry             interface{}
	RegistryExperimental interface{}
}

func (o *AttachSignatureOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *AttachSBOMOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *AttachSBOMOptions) MediaType() (interface{}, interface{}) {
	panic("stub")
}

func (o *AttachAttestationOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type AttestOptions struct {
	Key              interface{}
	Cert             interface{}
	CertChain        interface{}
	NoUpload         interface{}
	Recursive        interface{}
	Replace          interface{}
	SkipConfirmation interface{}
	TlogUpload       interface{}
	TSAServerURL     interface{}
	Rekor            interface{}
	Fulcio           interface{}
	OIDC             interface{}
	SecurityKey      interface{}
	Predicate        interface{}
	Registry         interface{}
}

func (o *AttestOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type AttestBlobOptions struct {
	Key                  interface{}
	Cert                 interface{}
	CertChain            interface{}
	SkipConfirmation     interface{}
	TlogUpload           interface{}
	TSAServerURL         interface{}
	RFC3161TimestampPath interface{}
	Hash                 interface{}
	Predicate            interface{}
	OutputSignature      interface{}
	OutputAttestation    interface{}
	OutputCertificate    interface{}
	BundlePath           interface{}
	Rekor                interface{}
	Fulcio               interface{}
	OIDC                 interface{}
	SecurityKey          interface{}
}

func (o *AttestBlobOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type CertVerifyOptions struct {
	Cert                         interface{}
	CertIdentity                 interface{}
	CertIdentityRegexp           interface{}
	CertOidcIssuer               interface{}
	CertOidcIssuerRegexp         interface{}
	CertGithubWorkflowTrigger    interface{}
	CertGithubWorkflowSha        interface{}
	CertGithubWorkflowName       interface{}
	CertGithubWorkflowRepository interface{}
	CertGithubWorkflowRef        interface{}
	CertChain                    interface{}
	SCT                          interface{}
	IgnoreSCT                    interface{}
}

func (o *CertVerifyOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *CertVerifyOptions) Identities() (interface{}, interface{}) {
	panic("stub")
}

type CleanType string

type CleanOptions struct {
	Registry  interface{}
	CleanType interface{}
	Force     interface{}
}

func defaultCleanType() interface{} {
	panic("stub")
}

func (c *CleanType) String() interface{} {
	panic("stub")
}

func (c *CleanType) Set(v interface{}) interface{} {
	panic("stub")
}

func (c *CleanType) Type() interface{} {
	panic("stub")
}

func (c *CleanOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type CopyOptions struct {
	SignatureOnly interface{}
	Force         interface{}
	Platform      interface{}
	Registry      interface{}
}

func (o *CopyOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type SBOMDownloadOptions struct{ Platform interface{} }

type AttestationDownloadOptions struct {
	PredicateType interface{}
	Platform      interface{}
}

func (o *SBOMDownloadOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *AttestationDownloadOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type EnvOptions struct {
	ShowDescriptions    interface{}
	ShowSensitiveValues interface{}
}

func (o *EnvOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type KeyParseError struct{}

type PubKeyParseError struct{}

func (e *KeyParseError) Error() interface{} {
	panic("stub")
}

func (e *PubKeyParseError) Error() interface{} {
	panic("stub")
}

func EnableExperimental() interface{} {
	panic("stub")
}

type FilesOptions struct{ Files interface{} }

func (o *FilesOptions) Parse() (interface{}, interface{}) {
	panic("stub")
}

func (o *FilesOptions) String() interface{} {
	panic("stub")
}

func (o *FilesOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func OneOf(args interface{}) interface{} {
	panic("stub")
}

func NOf(args interface{}) interface{} {
	panic("stub")
}

type FulcioOptions struct {
	URL                      interface{}
	IdentityToken            interface{}
	InsecureSkipFulcioVerify interface{}
}

func (o *FulcioOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type GenerateOptions struct {
	Embedme
	Registry interface{}
}

func (o *GenerateOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type GenerateKeyPairOptions struct {
	KMS             interface{}
	OutputKeyPrefix interface{}
}

func (o *GenerateKeyPairOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type ImportKeyPairOptions struct {
	Key             interface{}
	OutputKeyPrefix interface{}
}

func (o *ImportKeyPairOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type InitializeOptions struct {
	Mirror interface{}
	Root   interface{}
}

func (o *InitializeOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type KeyOpts struct {
	Sk                             interface{}
	Slot                           interface{}
	KeyRef                         interface{}
	FulcioURL                      interface{}
	RekorURL                       interface{}
	IDToken                        interface{}
	PassFunc                       interface{}
	OIDCIssuer                     interface{}
	OIDCClientID                   interface{}
	OIDCClientSecret               interface{}
	OIDCRedirectURL                interface{}
	OIDCDisableProviders           interface{}
	OIDCProvider                   interface{}
	BundlePath                     interface{}
	SkipConfirmation               interface{}
	TSAClientCACert                interface{}
	TSAClientCert                  interface{}
	TSAClientKey                   interface{}
	TSAServerName                  interface{}
	TSAServerURL                   interface{}
	RFC3161TimestampPath           interface{}
	TSACertChainPath               interface{}
	IssueCertificateForExistingKey interface{}
	FulcioAuthFlow                 interface{}
	InsecureSkipFulcioVerify       interface{}
}

type LoadOptions struct {
	Directory interface{}
	Registry  interface{}
}

func (o *LoadOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type OIDCOptions struct {
	Issuer                  interface{}
	ClientID                interface{}
	clientSecretFile        interface{}
	RedirectURL             interface{}
	Provider                interface{}
	DisableAmbientProviders interface{}
}

func (o *OIDCOptions) ClientSecret() (interface{}, interface{}) {
	panic("stub")
}

func (o *OIDCOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type Interface interface{}

type PIVToolSetManagementKeyOptions struct {
	OldKey    interface{}
	NewKey    interface{}
	RandomKey interface{}
}

type PIVToolSetPINOptions struct {
	OldPIN interface{}
	NewPIN interface{}
}

type PIVToolSetPUKOptions struct {
	OldPUK interface{}
	NewPUK interface{}
}

type PIVToolUnblockOptions struct {
	PUK    interface{}
	NewPIN interface{}
}

type PIVToolAttestationOptions struct {
	Output interface{}
	Slot   interface{}
}

type PIVToolGenerateKeyOptions struct {
	ManagementKey interface{}
	RandomKey     interface{}
	Slot          interface{}
	PINPolicy     interface{}
	TouchPolicy   interface{}
}

func (o *PIVToolSetManagementKeyOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *PIVToolSetPINOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *PIVToolSetPUKOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *PIVToolUnblockOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *PIVToolAttestationOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *PIVToolGenerateKeyOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type PKCS11ToolListTokensOptions struct{ ModulePath interface{} }

type PKCS11ToolListKeysUrisOptions struct {
	ModulePath interface{}
	SlotID     interface{}
	Pin        interface{}
}

func (o *PKCS11ToolListTokensOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *PKCS11ToolListKeysUrisOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type PredicateOptions struct{ Type interface{} }

type PredicateLocalOptions struct {
	Embedme
	Path interface{}
}

type PredicateRemoteOptions struct{ Embedme }

func (o *PredicateOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func ParsePredicateType(t interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (o *PredicateLocalOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *PredicateRemoteOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type PublicKeyOptions struct {
	Key         interface{}
	SecurityKey interface{}
	OutFile     interface{}
}

func (o *PublicKeyOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type ReferenceOptions struct{ TagPrefix interface{} }

func (o *ReferenceOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type RegistryExperimentalOptions struct{ RegistryReferrersMode interface{} }

type Keychain struct{}

type RegistryOptions struct {
	AllowInsecure      interface{}
	AllowHTTPRegistry  interface{}
	KubernetesKeychain interface{}
	RefOpts            interface{}
	Keychain           interface{}
	RegistryClientOpts interface{}
}

type RegistryReferrersMode string

func (o *RegistryOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *RegistryOptions) ClientOpts(ctx interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (o *RegistryOptions) NameOptions() interface{} {
	panic("stub")
}

func (o *RegistryOptions) GetRegistryClientOpts(ctx interface{}) interface{} {
	panic("stub")
}

func (e *RegistryReferrersMode) String() interface{} {
	panic("stub")
}

func (e *RegistryReferrersMode) Set(v interface{}) interface{} {
	panic("stub")
}

func (e *RegistryReferrersMode) Type() interface{} {
	panic("stub")
}

func (o *RegistryExperimentalOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type RekorOptions struct{ URL interface{} }

func (o *RekorOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type RootOptions struct {
	OutputFile interface{}
	Verbose    interface{}
	Timeout    interface{}
}

func (o *RootOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func BindViper(cmd interface{}, args interface{}) {
	panic("stub")
}

func callPersistentPreRun(cmd interface{}, args interface{}) {
	panic("stub")
}

func bindFlags(cmd interface{}, v interface{}) {
	panic("stub")
}

func flagToEnvVar(f interface{}) interface{} {
	panic("stub")
}

type SaveOptions struct{ Directory interface{} }

func (o *SaveOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type SecurityKeyOptions struct {
	Use  interface{}
	Slot interface{}
}

func (o *SecurityKeyOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type SignOptions struct {
	Key                   interface{}
	Cert                  interface{}
	CertChain             interface{}
	Upload                interface{}
	Output                interface{}
	OutputSignature       interface{}
	OutputPayload         interface{}
	OutputCertificate     interface{}
	PayloadPath           interface{}
	Recursive             interface{}
	Attachment            interface{}
	SkipConfirmation      interface{}
	TlogUpload            interface{}
	TSAClientCACert       interface{}
	TSAClientCert         interface{}
	TSAClientKey          interface{}
	TSAServerName         interface{}
	TSAServerURL          interface{}
	IssueCertificate      interface{}
	SignContainerIdentity interface{}
	Rekor                 interface{}
	Fulcio                interface{}
	OIDC                  interface{}
	SecurityKey           interface{}
	Embedme
	Registry             interface{}
	RegistryExperimental interface{}
}

func (o *SignOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type SignatureDigestOptions struct{ AlgorithmName interface{} }

func supportedSignatureAlgorithmNames() interface{} {
	panic("stub")
}

func (o *SignatureDigestOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *SignatureDigestOptions) HashAlgorithm() (interface{}, interface{}) {
	panic("stub")
}

type SignBlobOptions struct {
	Key                  interface{}
	Base64Output         interface{}
	Output               interface{}
	OutputSignature      interface{}
	OutputCertificate    interface{}
	SecurityKey          interface{}
	Fulcio               interface{}
	Rekor                interface{}
	OIDC                 interface{}
	Registry             interface{}
	BundlePath           interface{}
	SkipConfirmation     interface{}
	TlogUpload           interface{}
	TSAClientCACert      interface{}
	TSAClientCert        interface{}
	TSAClientKey         interface{}
	TSAServerName        interface{}
	TSAServerURL         interface{}
	RFC3161TimestampPath interface{}
	IssueCertificate     interface{}
}

func (o *SignBlobOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type TreeOptions struct {
	Registry  interface{}
	CleanType interface{}
}

func (c *TreeOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type TriangulateOptions struct {
	Type     interface{}
	Registry interface{}
}

func (o *TriangulateOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

type UploadBlobOptions struct {
	ContentType interface{}
	Files       interface{}
	Registry    interface{}
	Annotations interface{}
}

type UploadWASMOptions struct {
	File     interface{}
	Registry interface{}
}

func (o *UploadBlobOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *UploadWASMOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func UserAgent() interface{} {
	panic("stub")
}

type VerifyDockerfileOptions struct {
	Embedme
	BaseImageOnly interface{}
}

type VerifyBlobAttestationOptions struct {
	Key           interface{}
	SignaturePath interface{}
	BundlePath    interface{}
	Embedme
	CheckClaims          interface{}
	SecurityKey          interface{}
	CertVerify           interface{}
	Rekor                interface{}
	CommonVerifyOptions  interface{}
	RFC3161TimestampPath interface{}
}

type CommonVerifyOptions struct {
	Offline          interface{}
	TSACertChainPath interface{}
	IgnoreTlog       interface{}
	MaxWorkers       interface{}
}

type VerifyOptions struct {
	Key                 interface{}
	CheckClaims         interface{}
	Attachment          interface{}
	Output              interface{}
	SignatureRef        interface{}
	PayloadRef          interface{}
	LocalImage          interface{}
	CommonVerifyOptions interface{}
	SecurityKey         interface{}
	CertVerify          interface{}
	Rekor               interface{}
	Registry            interface{}
	SignatureDigest     interface{}
	Embedme
}

type VerifyAttestationOptions struct {
	Key                 interface{}
	CheckClaims         interface{}
	Output              interface{}
	CommonVerifyOptions interface{}
	SecurityKey         interface{}
	Rekor               interface{}
	CertVerify          interface{}
	Registry            interface{}
	Predicate           interface{}
	Policies            interface{}
	LocalImage          interface{}
}

type VerifyBlobOptions struct {
	Key                  interface{}
	Signature            interface{}
	BundlePath           interface{}
	SecurityKey          interface{}
	CertVerify           interface{}
	Rekor                interface{}
	CommonVerifyOptions  interface{}
	RFC3161TimestampPath interface{}
}

func (o *CommonVerifyOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *VerifyOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *VerifyAttestationOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *VerifyBlobOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *VerifyDockerfileOptions) AddFlags(cmd interface{}) {
	panic("stub")
}

func (o *VerifyBlobAttestationOptions) AddFlags(cmd interface{}) {
	panic("stub")
}
