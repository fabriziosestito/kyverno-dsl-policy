package errors

import _ "unsafe"

type Embedme interface{}

func WrapError(err interface{}) interface{} {
 panic("stub")
}

type CosignError struct {Message interface{}; Code interface{}}

func Error(cosignError interface{}) interface{} {
 panic("stub")
}

func (ce *CosignError) Error() interface{} {
 panic("stub")
}

func (ce *CosignError) ExitCode() interface{} {
 panic("stub")
}

func LookupExitCodeForError(err interface{}) interface{} {
 panic("stub")
}

func noMatchingSignatureError(err interface{}) interface{} {
 panic("stub")
}

func imageTagNotFoundError(err interface{}) interface{} {
 panic("stub")
}

func noSignaturesFoundError(err interface{}) interface{} {
 panic("stub")
}

func noCertificateFoundOnSignature(err interface{}) interface{} {
 panic("stub")
}

type ErrorCode struct {code interface{}; doc interface{}}

func GenerateExitCodeDocs(dir interface{}) interface{} {
 panic("stub")
}

func mustParse(fset interface{}, filename interface{}) interface{} {
 panic("stub")
}

