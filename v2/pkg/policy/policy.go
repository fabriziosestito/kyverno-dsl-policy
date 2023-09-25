package policy

import _ "unsafe"

type Embedme interface{}

func AttestationToPayloadJSON(_ interface{}, predicateType interface{}, verifiedAttestation interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

type EvaluationFailure struct {err interface{}}

func (e *EvaluationFailure) Error() interface{} {
 panic("stub")
}

func (e *EvaluationFailure) Unwrap() interface{} {
 panic("stub")
}

func EvaluatePolicyAgainstJSON(ctx interface{}, name, policyType interface{}, policyBody interface{}, jsonBytes interface{}) (warnings interface{}, errors interface{}) {
 panic("stub")
}

func evaluateCue(_ interface{}, attestation interface{}, evaluator interface{}) interface{} {
 panic("stub")
}

func evaluateRego(_ interface{}, attestation interface{}, evaluator interface{}) (warnings interface{}, errors interface{}) {
 panic("stub")
}

