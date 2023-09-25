package util

import _ "unsafe"

type Embedme interface{}

type CacheType string

type OnMemoryCache struct {TTL interface{}; data interface{}; mu interface{}}

type cachedObject struct {timestamp interface{}; object interface{}}

type Cache interface {}

type LocalFileCache struct {TTL interface{}; baseDir interface{}; mem interface{}}

func (c *OnMemoryCache) Set(key interface{}, value interface{}) interface{} {
 panic("stub")
}

func (c *OnMemoryCache) Get(key interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (c *OnMemoryCache) clearExpiredData() {
 panic("stub")
}

func (c *LocalFileCache) Set(key interface{}, value interface{}) interface{} {
 panic("stub")
}

func (c *LocalFileCache) Get(key interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (c *LocalFileCache) genFileNameFromKey(key interface{}) interface{} {
 panic("stub")
}

func (c *LocalFileCache) initBaseDir() interface{} {
 panic("stub")
}

func (c *LocalFileCache) baseDirExists() interface{} {
 panic("stub")
}

func (c *LocalFileCache) clearExpiredData() interface{} {
 panic("stub")
}

func (c *LocalFileCache) initMem() interface{} {
 panic("stub")
}

func GetCacheBaseDir() interface{} {
 panic("stub")
}

func IsLocalCacheEnabeld() interface{} {
 panic("stub")
}

func initGlobalCache() {
 panic("stub")
}

func SetCache(key interface{}, value interface{}) interface{} {
 panic("stub")
}

func GetCache(key interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetNameInfoFromCert(cert interface{}) interface{} {
 panic("stub")
}

func CmdExec(baseCmd interface{}, args interface{}) (interface{}, interface{}) {
 panic("stub")
}

type AnnotationWriter func(interface{}, interface{}) (interface{}, interface{})

type MutateOptions struct {AW interface{}; Annotations interface{}}

func TarGzCompress(src interface{}, buf interface{}, moList interface{}) interface{} {
 panic("stub")
}

func TarGzDecompress(src interface{}, dst interface{}) interface{} {
 panic("stub")
}

func copyDir(src interface{}, dst interface{}) interface{} {
 panic("stub")
}

func copyFile(src interface{}, dst interface{}) interface{} {
 panic("stub")
}

func IsDir(path interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetHomeDir() interface{} {
 panic("stub")
}

func LoadFileDataInEnvVar(envVarRef interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getSourceDigest(srcPath interface{}, moList interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GzipCompress(in interface{}) interface{} {
 panic("stub")
}

func GzipDecompress(in interface{}) interface{} {
 panic("stub")
}

func PullImage(resBundleRef interface{}, allowInsecure interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetImageDigest(resBundleRef interface{}, allowInsecure interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetBlob(layer interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GenerateConcatYAMLsFromImage(img interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetYAMLsInArtifact(blob interface{}) (interface{}, interface{}) {
 panic("stub")
}

func MatchPattern(pattern, value interface{}) interface{} {
 panic("stub")
}

func MatchSinglePattern(pattern, value interface{}) interface{} {
 panic("stub")
}

func ExactMatch(pattern, value interface{}) interface{} {
 panic("stub")
}

func ExactMatchWithPatternArray(value interface{}, patternArray interface{}) interface{} {
 panic("stub")
}

func GetUnionOfArrays(array1, array2 interface{}) interface{} {
 panic("stub")
}

func MatchPatternWithArray(pattern interface{}, valueArray interface{}) interface{} {
 panic("stub")
}

func MatchWithPatternArray(value interface{}, patternArray interface{}) interface{} {
 panic("stub")
}

func MatchBigInt(pattern interface{}, value interface{}) interface{} {
 panic("stub")
}

func pattern2BigInt(pattern interface{}) interface{} {
 panic("stub")
}

func SplitRule(rules interface{}) interface{} {
 panic("stub")
}

func SplitCommaSeparatedString(in interface{}) interface{} {
 panic("stub")
}

func init() {
 panic("stub")
}

type VersionInfo struct {Major interface{}; Minor interface{}; GitVersion interface{}; GitCommit interface{}; GitTreeState interface{}; BuildDate interface{}; GoVersion interface{}; Compiler interface{}; Platform interface{}}

func GetVersionInfo() interface{} {
 panic("stub")
}

func parseGitVersion(gv interface{}) (interface{}, interface{}) {
 panic("stub")
}

type candidateManifest struct {yaml interface{}; table interface{}; count interface{}; name interface{}}

type ResourceInfo struct {group interface{}; version interface{}; kind interface{}; name interface{}; namespace interface{}; raw interface{}}

func (ri ResourceInfo) Map() interface{} {
 panic("stub")
}

func FindYAMLsInDir(dirPath interface{}) (interface{}, interface{}) {
 panic("stub")
}

func LoadYAMLsInDirWithMutationOptions(dirPath interface{}, moList interface{}) (interface{}, interface{}) {
 panic("stub")
}

func FindManifestYAML(concatYamlBytes, objBytes interface{}, maxResourceManifestNum interface{}, ignoreFields interface{}) (interface{}, interface{}) {
 panic("stub")
}

func ManifestSearchByGVKNameNamespace(concatYamlBytes interface{}, apiVersion, kind, name, namespace interface{}) (interface{}, interface{}) {
 panic("stub")
}

func ManifestSearchByValue(concatYamlBytes, objBytes interface{}, maxResourceManifests interface{}, ignoreFields interface{}) (interface{}, interface{}) {
 panic("stub")
}

func extractKindMatchedManifests(concatYamlBytes interface{}, kind interface{}) interface{} {
 panic("stub")
}

func ConcatenateYAMLs(yamls interface{}) interface{} {
 panic("stub")
}

func IsConcatYAMLs(yaml interface{}) interface{} {
 panic("stub")
}

func SplitConcatYAMLs(yaml interface{}) interface{} {
 panic("stub")
}

func GetAnnotationsInYAML(yamlBytes interface{}) interface{} {
 panic("stub")
}

func parseResourceInfo(concatYamlBytes interface{}) interface{} {
 panic("stub")
}

func matchResourceInfo(msgInfos, reqInfos interface{}, useKeys interface{}) interface{} {
 panic("stub")
}

func isK8sResourceYAML(data interface{}) interface{} {
 panic("stub")
}

