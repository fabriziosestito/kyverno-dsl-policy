package k8smanifest

import _ "unsafe"

type Embedme interface{}

type MessageNotFoundError struct{ Embedme }

type K8sManifestError struct {
	Embedme
	defaultMessage interface{}
}

type SignatureNotFoundError struct{ Embedme }

type SignatureVerificationError struct{ Embedme }

func (e *K8sManifestError) Error() interface{} {
	panic("stub")
}

func NewSignatureNotFoundError(err interface{}) interface{} {
	panic("stub")
}

func NewMessageNotFoundError(err interface{}) interface{} {
	panic("stub")
}

func NewSignatureVerificationError(err interface{}) interface{} {
	panic("stub")
}

func IsSignatureNotFoundError(err interface{}) interface{} {
	panic("stub")
}

func IsMessageNotFoundError(err interface{}) interface{} {
	panic("stub")
}

func IsSignatureVerificationError(err interface{}) interface{} {
	panic("stub")
}

type SignerList []string

type ObjectFieldBinding struct {
	Fields  interface{}
	Objects interface{}
}

type cosignSignOption struct {
	RekorURL      interface{}
	NoTlogUpload  interface{}
	AllowInsecure interface{}
	Force         interface{}
}

type cosignVerifyOption struct {
	Certificate      interface{}
	CertificateChain interface{}
	RekorURL         interface{}
	OIDCIssuer       interface{}
	RootCerts        interface{}
	AllowInsecure    interface{}
}

type ObjectReference struct {
	Group     interface{}
	Version   interface{}
	Kind      interface{}
	Name      interface{}
	Namespace interface{}
}

type VerifyResourceOption struct {
	Embedme
	SkipObjects           interface{}
	Provenance            interface{}
	DisableDryRun         interface{}
	CheckDryRunForApply   interface{}
	CheckMutatingResource interface{}
	DryRunNamespace       interface{}
}

type verifyOption struct {
	IgnoreFields                interface{}
	Signers                     interface{}
	MaxResourceManifestNum      interface{}
	KeyPath                     interface{}
	ResourceBundleRef           interface{}
	SignatureResourceRef        interface{}
	ProvenanceResourceRef       interface{}
	UseCache                    interface{}
	CacheDir                    interface{}
	annotationKeyToIgnoreFields interface{}
}

type SignOption struct {
	Embedme
	KeyPath           interface{}
	ResourceBundleRef interface{}
	CertPath          interface{}
	Output            interface{}
	UpdateAnnotation  interface{}
	ImageAnnotations  interface{}
	PassFunc          interface{}
	ApplySigConfigMap interface{}
	Tarball           interface{}
	AppendSignature   interface{}
}

type ObjectFieldBindingList []ObjectFieldBinding

type ObjectReferenceList []ObjectReference

type ObjectUserBinding struct {
	Users   interface{}
	Objects interface{}
}

type VerifyManifestOption struct{ Embedme }

type commonOption struct{ Embedme }

type AnnotationConfig struct {
	AnnotationKeyDomain       interface{}
	ResourceBundleRefBaseName interface{}
	SignatureBaseName         interface{}
	CertificateBaseName       interface{}
	MessageBaseName           interface{}
	BundleBaseName            interface{}
}

func (o *VerifyResourceOption) SetAnnotationIgnoreFields() {
	panic("stub")
}

func (o *VerifyManifestOption) SetAnnotationIgnoreFields() {
	panic("stub")
}

func (o verifyOption) setAnnotationKeyToIgnoreField(annotationConfig interface{}) interface{} {
	panic("stub")
}

func (o verifyOption) isAnnotationKeyAlreadySetToIgnoreFields() interface{} {
	panic("stub")
}

func (c AnnotationConfig) MessageAnnotationKey() interface{} {
	panic("stub")
}

func (c AnnotationConfig) SignatureAnnotationKey(i interface{}) interface{} {
	panic("stub")
}

func (c AnnotationConfig) CertificateAnnotationKey(i interface{}) interface{} {
	panic("stub")
}

func (c AnnotationConfig) BundleAnnotationKey(i interface{}) interface{} {
	panic("stub")
}

func (c AnnotationConfig) ResourceBundleRefAnnotationKey() interface{} {
	panic("stub")
}

func (c AnnotationConfig) annotationKey(keyType interface{}, i interface{}) interface{} {
	panic("stub")
}

func (c AnnotationConfig) AnnotationKeyMap(i interface{}) interface{} {
	panic("stub")
}

func (c AnnotationConfig) GetAllSignatureSets(annotations interface{}) interface{} {
	panic("stub")
}

func (c AnnotationConfig) AnnotationKeyMask() interface{} {
	panic("stub")
}

func (c AnnotationConfig) AnnotationKeyIgnoreField() interface{} {
	panic("stub")
}

func firstNonEmpty(strArray interface{}) interface{} {
	panic("stub")
}

func ObjectToReference(obj interface{}) interface{} {
	panic("stub")
}

func (l ObjectReferenceList) Match(obj interface{}) interface{} {
	panic("stub")
}

func (r ObjectReference) Match(obj interface{}) interface{} {
	panic("stub")
}

func (r ObjectReference) Equal(r2 interface{}) interface{} {
	panic("stub")
}

func (l ObjectFieldBindingList) Match(obj interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (f ObjectFieldBinding) Match(obj interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (l SignerList) Match(signerName interface{}) interface{} {
	panic("stub")
}

func LoadVerifyManifestConfig(fpath interface{}) (interface{}, interface{}) {
	panic("stub")
}

func LoadVerifyResourceConfig(fpath interface{}) (interface{}, interface{}) {
	panic("stub")
}

func LoadVerifyResourceConfigFromResource(configPath, configField interface{}) (interface{}, interface{}) {
	panic("stub")
}

func GetConfigResource(configPath interface{}) (interface{}, interface{}) {
	panic("stub")
}

func GetMatchConditionFromConfigResource(configPath, matchField, inScopeObjectField interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func parseConfigObj(configObj interface{}, configField interface{}) (interface{}, interface{}) {
	panic("stub")
}

func getMatchConditionInConstraint(configObj interface{}, matchField interface{}) (interface{}, interface{}) {
	panic("stub")
}

func parseInScopeObjectInConstraint(configObj interface{}, inScopeObjectField interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (vo *VerifyResourceOption) AddDefaultConfig(defaultConfig interface{}) interface{} {
	panic("stub")
}

func LoadDefaultConfig() interface{} {
	panic("stub")
}

func AddDefaultConfig(vo interface{}) interface{} {
	panic("stub")
}

type resourceName struct {
	PodName       interface{}
	ContainerName interface{}
}

type RecursiveImageProvenanceGetter struct {
	object                        interface{}
	manifestResourceBundleRef     interface{}
	manifestProvenanceResourceRef interface{}
	cacheEnabled                  interface{}
	allowInsecure                 interface{}
}

type ProvenanceMaterial struct {
	URI    interface{}
	Digest interface{}
}

type Provenance struct {
	ResourceName         interface{}
	RawAttestation       interface{}
	RawSBOM              interface{}
	Artifact             interface{}
	ArtifactType         interface{}
	Hash                 interface{}
	AttestationLogIndex  interface{}
	AttestationMaterials interface{}
	SBOMRef              interface{}
	ConfigMapRef         interface{}
}

type ResourceProvenanceGetter struct{ resourceRefString interface{} }

type NotImplementedProvenanceGetter struct{}

type ProvenanceGetter interface{}

type DigestSet map[string]string

type ImageProvenanceGetter struct {
	resBundleRef  interface{}
	imageHash     interface{}
	cacheEnabled  interface{}
	allowInsecure interface{}
}

type rekorCLIGetCmdOutput struct {
	Attestation     interface{}
	AttestationType interface{}
	Body            interface{}
	LogIndex        interface{}
	IntegratedTime  interface{}
	UUID            interface{}
	LogID           interface{}
}

type ArtifactType string

func NewProvenanceGetter(obj interface{}, sigRef, imageHash, provResRef interface{}, allowInsecure interface{}) interface{} {
	panic("stub")
}

func (g *RecursiveImageProvenanceGetter) Get() (interface{}, interface{}) {
	panic("stub")
}

func (g *RecursiveImageProvenanceGetter) imageDigest(resBundleRef interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (g *ImageProvenanceGetter) Get() (interface{}, interface{}) {
	panic("stub")
}

func (g *ImageProvenanceGetter) getAttestation() (interface{}, interface{}, interface{}) {
	panic("stub")
}

func (g *ImageProvenanceGetter) getSBOM() (interface{}, interface{}) {
	panic("stub")
}

func (g *ImageProvenanceGetter) getSBOMRef(resBundleRef interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (g *ResourceProvenanceGetter) Get() (interface{}, interface{}) {
	panic("stub")
}

func (g *ResourceProvenanceGetter) getProvenanceInSingleConfigMap(singleCMRef interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (g *NotImplementedProvenanceGetter) Get() (interface{}, interface{}) {
	panic("stub")
}

func getAttestationEntry(hash interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func parseEntry(uuid interface{}, e interface{}) (interface{}, interface{}) {
	panic("stub")
}

func ParseAttestation(attestationStr interface{}) (interface{}, interface{}, interface{}, interface{}) {
	panic("stub")
}

func GenerateIntotoAttestationCurlCommand(logIndex interface{}) interface{} {
	panic("stub")
}

func GenerateIntotoAttestationKubectlCommand(resourceRef interface{}) interface{} {
	panic("stub")
}

func GenerateSBOMDownloadCommand(resBundleRef interface{}) interface{} {
	panic("stub")
}

func GenerateSBOMKubectlCommand(resourceRef interface{}) interface{} {
	panic("stub")
}

type CosignSignConfig struct {
	RekorURL      interface{}
	NoTlogUpload  interface{}
	AllowInsecure interface{}
	Force         interface{}
}

type BlobSigner struct {
	AnnotationConfig   interface{}
	createSigConfigMap interface{}
	appendSig          interface{}
	doApply            interface{}
	tarball            interface{}
	prikeyPath         interface{}
	certPath           interface{}
	passFunc           interface{}
	Embedme
}

type Signer interface{}

type ImageSigner struct {
	AnnotationConfig interface{}
	tarball          interface{}
	resBundleRef     interface{}
	prikeyPath       interface{}
	certPath         interface{}
	passFunc         interface{}
	Embedme
}

func Sign(inputDir interface{}, so interface{}) (interface{}, interface{}) {
	panic("stub")
}

func NewSigner(resBundleRef, keyPath, certPath, output interface{}, appendSig, doApply, tarball interface{}, cosignSignConfig interface{}, AnnotationConfig interface{}, pf interface{}) interface{} {
	panic("stub")
}

func (s *ImageSigner) Sign(inputDir, output interface{}, imageAnnotations interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (s *BlobSigner) Sign(inputDir, output interface{}, imageAnnotations interface{}) (interface{}, interface{}) {
	panic("stub")
}

func makeMessageYAML(inputDir interface{}, outBuffer interface{}, moList interface{}) interface{} {
	panic("stub")
}

func uploadFileToRegistry(inputData interface{}, resBundleRef interface{}, allowInsecure interface{}, imageAnnotations interface{}) interface{} {
	panic("stub")
}

func generateSignatureConfigMap(sigResRef interface{}, sigMaps interface{}) (interface{}, interface{}) {
	panic("stub")
}

func generateSignedYAMLManifest(inputDir, resBundleRef interface{}, sigMaps interface{}, appendSig interface{}, imageAnnotations interface{}, AnnotationConfig interface{}) (interface{}, interface{}) {
	panic("stub")
}

func embedAnnotation(yamlBytes interface{}, annotationMap interface{}) (interface{}, interface{}) {
	panic("stub")
}

func getMutationOptionForClean(annoCfg interface{}) interface{} {
	panic("stub")
}

func removeSignatureAnnotation(yamlBytes interface{}, annotationMap interface{}) (interface{}, interface{}) {
	panic("stub")
}

func applySignatureConfigMap(configMapRef interface{}, newCM interface{}) (interface{}, interface{}) {
	panic("stub")
}

func K8sResourceRef2FileName(resRef interface{}) interface{} {
	panic("stub")
}

type SignatureVerifier interface{}

type ImageSignatureVerifier struct {
	resBundleRef         interface{}
	onMemoryCacheEnabled interface{}
	annotationConfig     interface{}
	identityList         interface{}
	Embedme
}

type ManifestFetcher interface{}

type verificationIdentity struct {
	path interface{}
	name interface{}
}

type CosignVerifyConfig struct {
	CertRef       interface{}
	CertChain     interface{}
	RekorURL      interface{}
	OIDCIssuer    interface{}
	RootCerts     interface{}
	AllowInsecure interface{}
}

type BlobManifestFetcher struct {
	AnnotationConfig       interface{}
	resourceRefString      interface{}
	ignoreFields           interface{}
	maxResourceManifestNum interface{}
}

type VerifyResult struct {
	Verified interface{}
	Signer   interface{}
	Diff     interface{}
}

type BlobSignatureVerifier struct {
	annotations      interface{}
	resourceRef      interface{}
	annotationConfig interface{}
	identityList     interface{}
	Embedme
}

type ImageManifestFetcher struct {
	resBundleRefString     interface{}
	AnnotationConfig       interface{}
	ignoreFields           interface{}
	maxResourceManifestNum interface{}
	cacheEnabled           interface{}
	allowInsecure          interface{}
}

func NewSignatureVerifier(objYAMLBytes interface{}, sigRef interface{}, pubkeyPath interface{}, signers interface{}, cosignVerifyConfig interface{}, annotationConfig interface{}) interface{} {
	panic("stub")
}

func (v *ImageSignatureVerifier) Verify() (interface{}, interface{}, interface{}, interface{}) {
	panic("stub")
}

func (v *ImageSignatureVerifier) getResultFromCache(resBundleRef, pubkey interface{}) (interface{}, interface{}, interface{}, interface{}, interface{}) {
	panic("stub")
}

func (v *ImageSignatureVerifier) setResultToCache(resBundleRef, pubkey interface{}, verified interface{}, signerName interface{}, signedTimestamp interface{}, err interface{}) {
	panic("stub")
}

func (v *BlobSignatureVerifier) Verify() (interface{}, interface{}, interface{}, interface{}) {
	panic("stub")
}

func (v *BlobSignatureVerifier) getSignatureSets() (interface{}, interface{}) {
	panic("stub")
}

func NewManifestFetcher(resBundleRef, resourceRef interface{}, annotationConfig interface{}, ignoreFields interface{}, maxResourceManifestNum interface{}, allowInsecure interface{}) interface{} {
	panic("stub")
}

func (f *ImageManifestFetcher) Fetch(objYAMLBytes interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func (f *ImageManifestFetcher) fetchManifestInSingleImage(singleResourceBundleRef interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (f *ImageManifestFetcher) FetchAll() (interface{}, interface{}) {
	panic("stub")
}

func (f *ImageManifestFetcher) getConcatYAMLFromResourceBundleRef(resBundleRef interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (f *ImageManifestFetcher) getManifestFromCache(resBundleRef interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func (f *ImageManifestFetcher) setManifestToCache(resBundleRef interface{}, concatYAMLbytes interface{}, err interface{}) {
	panic("stub")
}

func (f *BlobManifestFetcher) Fetch(objYAMLBytes interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func (f *BlobManifestFetcher) fetchManifestFromResource(objYAMLBytes interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func (f *BlobManifestFetcher) fetchManifestInSingleConfigMap(singleCMRef interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (r *VerifyResult) String() interface{} {
	panic("stub")
}

func GetConfigMapFromK8sObjectRef(objRef interface{}) (interface{}, interface{}) {
	panic("stub")
}

func VerifyManifest(objManifest interface{}, vo interface{}) (interface{}, interface{}) {
	panic("stub")
}

func matchManifest(inputManifestBytes, foundManifestBytes interface{}, ignoreFields interface{}, AnnotationConfig interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

type VerifyResourceResult struct {
	Verified        interface{}
	InScope         interface{}
	Signer          interface{}
	SignedTime      interface{}
	SigRef          interface{}
	Diff            interface{}
	ContainerImages interface{}
	Provenances     interface{}
}

func (r *VerifyResourceResult) String() interface{} {
	panic("stub")
}

func VerifyResource(obj interface{}, vo interface{}) (interface{}, interface{}) {
	panic("stub")
}

func matchResourceWithManifest(obj interface{}, foundManifestBytes interface{}, ignoreFields interface{}, dryRunNamespace interface{}, disableDryRun, checkDryRunForApply, checkMutatingResource interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func directMatch(messageYAMLBytes, resourceJSONBytes interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func dryrunCreateMatch(messageYAMLBytes, resourceJSONBytes interface{}, clusterScope, isCRD interface{}, dryRunNamespace interface{}) (interface{}, interface{}, interface{}, interface{}) {
	panic("stub")
}

func dryrunApplyMatch(messageYAMLBytes, resourceJSONBytes interface{}, clusterScope, isCRD interface{}, dryRunNamespace interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func inclusionMatch(messageYAMLBytes, resourceJSONBytes, dryRunBytes interface{}, clusterScope, isCRD, disableDryRun interface{}) (interface{}, interface{}, interface{}) {
	panic("stub")
}

func removeCanonicalizedImageDiff(diff interface{}) interface{} {
	panic("stub")
}

func inverseDiff(diff interface{}) interface{} {
	panic("stub")
}

func getTime(tstamp interface{}) interface{} {
	panic("stub")
}
