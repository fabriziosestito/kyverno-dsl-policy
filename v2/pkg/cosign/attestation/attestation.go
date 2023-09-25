package attestation

import _ "unsafe"

type Embedme interface{}

type CosignVulnStatement struct {Embedme; Predicate interface{}}

type Invocation struct {Parameters interface{}; URI interface{}; EventID interface{}; BuilderID interface{}}

type DB struct {URI interface{}; Version interface{}}

type CosignVulnPredicate struct {Invocation interface{}; Scanner interface{}; Metadata interface{}}

type Scanner struct {URI interface{}; Version interface{}; DB interface{}; Result interface{}}

type GenerateOpts struct {Predicate interface{}; Type interface{}; Digest interface{}; Repo interface{}; Time interface{}}

type CosignPredicate struct {Data interface{}; Timestamp interface{}}

type Metadata struct {ScanStartedOn interface{}; ScanFinishedOn interface{}}

func GenerateStatement(opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func generateVulnStatement(predicate interface{}, digest interface{}, repo interface{}) (interface{}, interface{}) {
 panic("stub")
}

func timestamp(opts interface{}) interface{} {
 panic("stub")
}

func customType(opts interface{}) interface{} {
 panic("stub")
}

func generateStatementHeader(digest, repo, predicateType interface{}) interface{} {
 panic("stub")
}

func generateCustomStatement(rawPayload interface{}, customType, digest, repo, timestamp interface{}) (interface{}, interface{}) {
 panic("stub")
}

func generateCustomPredicate(rawPayload interface{}, customType, timestamp interface{}) (interface{}, interface{}) {
 panic("stub")
}

func generateSLSAProvenanceStatementSLSA02(rawPayload interface{}, digest interface{}, repo interface{}) (interface{}, interface{}) {
 panic("stub")
}

func generateSLSAProvenanceStatementSLSA1(rawPayload interface{}, digest interface{}, repo interface{}) (interface{}, interface{}) {
 panic("stub")
}

func generateLinkStatement(rawPayload interface{}, digest interface{}, repo interface{}) (interface{}, interface{}) {
 panic("stub")
}

func generateSPDXStatement(rawPayload interface{}, digest interface{}, repo interface{}, parseJSON interface{}) (interface{}, interface{}) {
 panic("stub")
}

func generateCycloneDXStatement(rawPayload interface{}, digest interface{}, repo interface{}) (interface{}, interface{}) {
 panic("stub")
}

func checkRequiredJSONFields(rawPayload interface{}, typ interface{}) interface{} {
 panic("stub")
}

