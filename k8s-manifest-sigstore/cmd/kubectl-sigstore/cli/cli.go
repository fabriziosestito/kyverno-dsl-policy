package cli

import _ "unsafe"

type Embedme interface{}

func NewCmdApplyAfterVerify() interface{} {
 panic("stub")
}

func applyAfterVerify(filename, resBundleRef, keyPath, configPath interface{}) interface{} {
 panic("stub")
}

type KubectlOptions struct {ConfigFlags interface{}; PrintFlags interface{}; GetOptions interface{}; ApplyOptions interface{}; fieldManagerForApply interface{}}

func (o *KubectlOptions) SetKubeConfig(fpath, namespace interface{}) {
 panic("stub")
}

func (o *KubectlOptions) InitGet(cmd interface{}) interface{} {
 panic("stub")
}

func (o *KubectlOptions) Get(args interface{}, overrideNamespace interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (o *KubectlOptions) InitApply(cmd interface{}, filename interface{}) interface{} {
 panic("stub")
}

func (o *KubectlOptions) Apply(filename interface{}) interface{} {
 panic("stub")
}

func NewCmdManifestBuild() interface{} {
 panic("stub")
}

func buildManifest(baseDir, outputPath, provenancePath, resBundleRef, keyPath interface{}, kustomizeMode, signProvenance interface{}) interface{} {
 panic("stub")
}

func init() {
 panic("stub")
}

func NewCmdSign() interface{} {
 panic("stub")
}

func sign(inputDir, resBundleRef, keyPath, rekorURL, output interface{}, appendSignature, applySignatureConfigMap, updateAnnotation, tarball, allowInsecure, noTlogUpload, force interface{}, annotations interface{}) interface{} {
 panic("stub")
}

func parseAnnotations(annotations interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewCmdVerify() interface{} {
 panic("stub")
}

func verify(filename, resBundleRef, keyPath, configPath, certRef, certChain, rekorURL, oidcIssuer interface{}, allowInsecure interface{}) interface{} {
 panic("stub")
}

type imageToBeUsed struct {Embedme; imageType interface{}}

type resourceResult struct {Object interface{}; Result interface{}; Error interface{}}

type provenanceSummary struct {Total interface{}; Artifacts interface{}}

type summary struct {Total interface{}; Valid interface{}; Invalid interface{}}

type manifestResult struct {Name interface{}; Signer interface{}; SignedTime interface{}}

type provenanceResult struct {Summary interface{}; Items interface{}}

type VerifyResourceResult struct {Embedme; Summary interface{}; Manifests interface{}; Resources interface{}; Provenance interface{}}

func NewCmdVerifyResource() interface{} {
 panic("stub")
}

func verifyResource(yamls interface{}, kubeGetArgs interface{}, resBundleRef, sigResRef, keyPath, configPath, configField, configType interface{}, disableDefaultConfig, provenance, allowInsecure interface{}, provResRef, certRef, certChain, rekorURL, oidcIssuer, outputFormat interface{}, concurrencyNum interface{}) (interface{}, interface{}) {
 panic("stub")
}

func readManifestYAMLFile(fpath interface{}) (interface{}, interface{}) {
 panic("stub")
}

func readStdinAsYAMLs() (interface{}, interface{}) {
 panic("stub")
}

func getObjsFromManifests(yamls interface{}, ignoreFieldConfig interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getObjsByConstraint(constraintRef, matchField, inscopeField interface{}, concurrencyNum interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getObjsByConstraintWithCache(constraintRef, matchField, inscopeField interface{}, concurrencyNum interface{}) (interface{}, interface{}) {
 panic("stub")
}

func makeResultTable(result interface{}, provenanceEnabled interface{}) interface{} {
 panic("stub")
}

func makeSummaryResultTable(result interface{}) interface{} {
 panic("stub")
}

func makeManifestResultTable(result interface{}, provenanceEnabled interface{}) interface{} {
 panic("stub")
}

func makeResourceResultTable(result interface{}, provenanceEnabled interface{}) interface{} {
 panic("stub")
}

func makeProvenanceResultTable(result interface{}) interface{} {
 panic("stub")
}

func getAge(t interface{}) interface{} {
 panic("stub")
}

func (r resourceResult) MarshalJSON() (interface{}, interface{}) {
 panic("stub")
}

func (r resourceResult) MarshalYAML() (interface{}, interface{}) {
 panic("stub")
}

func NewVerifyResourceResult(results interface{}, provenanceEnabled interface{}) interface{} {
 panic("stub")
}

func obj2ref(obj interface{}) interface{} {
 panic("stub")
}

func getConfigPathFromConfigFlags(path, ctype, kind, name, namespace, field interface{}) (interface{}, interface{}) {
 panic("stub")
}

func objsToConcatYAML(objs interface{}) interface{} {
 panic("stub")
}

func validateConfigMapRef(cmRefString interface{}) interface{} {
 panic("stub")
}

func getAllImagesToBeUsed(resBundleRef interface{}, objs interface{}, annotationConfig interface{}, provenanceEnabled interface{}) interface{} {
 panic("stub")
}

func NewCmdVersion() interface{} {
 panic("stub")
}

func version() interface{} {
 panic("stub")
}

