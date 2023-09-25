package attach

import _ "unsafe"

type Embedme interface{}

func AttestationCmd(ctx interface{}, regOpts interface{}, signedPayloads interface{}, imageRef interface{}) interface{} {
 panic("stub")
}

func attachAttestation(ctx interface{}, remoteOpts interface{}, signedPayload, imageRef interface{}, nameOpts interface{}) interface{} {
 panic("stub")
}

func SBOMCmd(ctx interface{}, regOpts interface{}, regExpOpts interface{}, sbomRef interface{}, sbomType interface{}, imageRef interface{}) interface{} {
 panic("stub")
}

func sbomCmdOCIExperimental(ctx interface{}, regOpts interface{}, sbomRef interface{}, sbomType interface{}, imageRef interface{}) interface{} {
 panic("stub")
}

func sbomBytes(sbomRef interface{}) (interface{}, interface{}) {
 panic("stub")
}

type SignatureArgType uint8

func SignatureCmd(ctx interface{}, regOpts interface{}, sigRef, payloadRef, certRef, certChainRef, timeStampedSigRef, imageRef interface{}) interface{} {
 panic("stub")
}

func signatureBytes(sigRef interface{}) (interface{}, interface{}) {
 panic("stub")
}

func signatureType(sigRef interface{}) interface{} {
 panic("stub")
}

