package rego

import _ "unsafe"

type Embedme interface{}

type CosignRuleResult struct {Warning interface{}; Error interface{}; Result interface{}}

func ValidateJSON(jsonBody interface{}, entrypoints interface{}) interface{} {
 panic("stub")
}

func ValidateJSONWithModuleInput(jsonBody interface{}, moduleInput interface{}) (warnings interface{}, errors interface{}) {
 panic("stub")
}

func evaluateRegoEvalMapResult(query interface{}, response interface{}) (warning interface{}, error interface{}) {
 panic("stub")
}

