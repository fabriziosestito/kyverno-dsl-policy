package cli

import _ "unsafe"

type Embedme interface{}

func Attach() interface{} {
 panic("stub")
}

func attachSignature() interface{} {
 panic("stub")
}

func attachSBOM() interface{} {
 panic("stub")
}

func attachAttestation() interface{} {
 panic("stub")
}

func Attest() interface{} {
 panic("stub")
}

func AttestBlob() interface{} {
 panic("stub")
}

func Clean() interface{} {
 panic("stub")
}

func CleanCmd(ctx interface{}, regOpts interface{}, cleanType interface{}, imageRef interface{}, force interface{}) interface{} {
 panic("stub")
}

func prompt(cleanType interface{}) interface{} {
 panic("stub")
}

func normalizeCertificateFlags(_ interface{}, name interface{}) interface{} {
 panic("stub")
}

func New() interface{} {
 panic("stub")
}

func Completion() interface{} {
 panic("stub")
}

func Copy() interface{} {
 panic("stub")
}

func Dockerfile() interface{} {
 panic("stub")
}

func dockerfileVerify() interface{} {
 panic("stub")
}

func Download() interface{} {
 panic("stub")
}

func downloadSignature() interface{} {
 panic("stub")
}

func downloadSBOM() interface{} {
 panic("stub")
}

func downloadAttestation() interface{} {
 panic("stub")
}

type envGetter func(interface{}) interface{}

type environGetter func() interface{}

func Env() interface{} {
 panic("stub")
}

func getEnv() interface{} {
 panic("stub")
}

func getEnviron() interface{} {
 panic("stub")
}

func printEnv(envVars interface{}, envGet interface{}, environGet interface{}, showDescription, showSensitive interface{}) {
 panic("stub")
}

func sortEnvKeys(envVars interface{}) interface{} {
 panic("stub")
}

func Generate() interface{} {
 panic("stub")
}

func GenerateKeyPair() interface{} {
 panic("stub")
}

func ImportKeyPair() interface{} {
 panic("stub")
}

func Initialize() interface{} {
 panic("stub")
}

func Load() interface{} {
 panic("stub")
}

func LoadCmd(ctx interface{}, opts interface{}, imageRef interface{}) interface{} {
 panic("stub")
}

func Manifest() interface{} {
 panic("stub")
}

func manifestVerify() interface{} {
 panic("stub")
}

func PIVTool() interface{} {
 panic("stub")
}

func PKCS11Tool() interface{} {
 panic("stub")
}

func PublicKey() interface{} {
 panic("stub")
}

func Save() interface{} {
 panic("stub")
}

func SaveCmd(_ interface{}, opts interface{}, imageRef interface{}) interface{} {
 panic("stub")
}

func Sign() interface{} {
 panic("stub")
}

func SignBlob() interface{} {
 panic("stub")
}

func Tree() interface{} {
 panic("stub")
}

func TreeCmd(ctx interface{}, regOpts interface{}, imageRef interface{}) interface{} {
 panic("stub")
}

func printLayers(layers interface{}) interface{} {
 panic("stub")
}

func Triangulate() interface{} {
 panic("stub")
}

func Upload() interface{} {
 panic("stub")
}

func uploadBlob() interface{} {
 panic("stub")
}

func uploadWASM() interface{} {
 panic("stub")
}

func Verify() interface{} {
 panic("stub")
}

func VerifyAttestation() interface{} {
 panic("stub")
}

func VerifyBlob() interface{} {
 panic("stub")
}

func VerifyBlobAttestation() interface{} {
 panic("stub")
}

