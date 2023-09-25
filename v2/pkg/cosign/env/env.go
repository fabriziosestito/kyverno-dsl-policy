package env

import _ "unsafe"

type Embedme interface{}

type Variable string

type VariableOpts struct {Description interface{}; Expects interface{}; Sensitive interface{}; External interface{}}

func (v Variable) String() interface{} {
 panic("stub")
}

func EnvironmentVariables() interface{} {
 panic("stub")
}

func mustRegisterEnv(name interface{}) {
 panic("stub")
}

func Getenv(name interface{}) interface{} {
 panic("stub")
}

func LookupEnv(name interface{}) (interface{}, interface{}) {
 panic("stub")
}

