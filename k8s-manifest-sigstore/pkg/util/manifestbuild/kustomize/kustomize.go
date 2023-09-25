package kustomize

import _ "unsafe"

type Embedme interface{}

type GitRepoResult struct {RootDir interface{}; URL interface{}; Revision interface{}; CommitID interface{}; Path interface{}}

type FileInfo struct {Name interface{}; Hash interface{}}

type KustomizationResource struct {GitRepo interface{}; File interface{}}

func LoadKustomization(fpath, baseDir, gitURL, gitRevision interface{}, inRemoteRepo interface{}) (interface{}, interface{}) {
 panic("stub")
}

func Sha256Hash(fpath interface{}) (interface{}, interface{}) {
 panic("stub")
}

func prepareBaseDirForRemoteRepository(url interface{}) (interface{}, interface{}) {
 panic("stub")
}

func parseGitURLinKustomization(urlInKustomization interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func CmdExec(baseCmd, dir interface{}, args interface{}) (interface{}, interface{}) {
 panic("stub")
}

func KustomizeExec(dir interface{}, args interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GitExec(dir interface{}, args interface{}) (interface{}, interface{}) {
 panic("stub")
}

func IsRepositoryResource(path interface{}) interface{} {
 panic("stub")
}

func IsFileResource(path interface{}) interface{} {
 panic("stub")
}

func IsFile(name interface{}) (interface{}, interface{}) {
 panic("stub")
}

func IsDir(name interface{}) (interface{}, interface{}) {
 panic("stub")
}

func FileExists(fpath interface{}) interface{} {
 panic("stub")
}

func parseGitUrl(n interface{}) (host interface{}, orgRepo interface{}, path interface{}, gitRef interface{}, gitSuff interface{}) {
 panic("stub")
}

func peelQuery(arg interface{}) (interface{}, interface{}) {
 panic("stub")
}

func parseHostSpec(n interface{}) (interface{}, interface{}) {
 panic("stub")
}

func normalizeGitHostSpec(host interface{}) interface{} {
 panic("stub")
}

type IntotoSigner struct {key interface{}; keyID interface{}}

func GenerateProvenance(artifactName, digest, kustomizeBase interface{}, startTime, finishTime interface{}, recipeCmd interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GenerateAttestation(provPath, privKeyPath interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetDigestOfArtifact(artifactPath interface{}) (interface{}, interface{}) {
 panic("stub")
}

func OverwriteArtifactInProvenance(provPath, overwriteArtifact interface{}) (interface{}, interface{}) {
 panic("stub")
}

func generateMaterialsFromKustomization(kustomizeBase interface{}) (interface{}, interface{}) {
 panic("stub")
}

func checkRepoInfoOfKustomizeBase(kustomizeBase interface{}) (interface{}, interface{}, interface{}, interface{}) {
 panic("stub")
}

func resourceToMaterial(kr interface{}) interface{} {
 panic("stub")
}

func GetImageDigest(resBundleRef interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (it *IntotoSigner) Sign(ctx interface{}, data interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (it *IntotoSigner) Verify(ctx interface{}, data, sig interface{}) interface{} {
 panic("stub")
}

func (es *IntotoSigner) KeyID() (interface{}, interface{}) {
 panic("stub")
}

func (es *IntotoSigner) Public() interface{} {
 panic("stub")
}

