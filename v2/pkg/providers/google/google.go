package google

import _ "unsafe"

type Embedme interface{}

type googleImpersonate struct {}

type googleWorkloadIdentity struct {}

func init() {
 panic("stub")
}

func (gwi *googleWorkloadIdentity) Enabled(ctx interface{}) interface{} {
 panic("stub")
}

func (gwi *googleWorkloadIdentity) Provide(ctx interface{}, audience interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (gi *googleImpersonate) Enabled(_ interface{}) interface{} {
 panic("stub")
}

func (gi *googleImpersonate) Provide(ctx interface{}, audience interface{}) (interface{}, interface{}) {
 panic("stub")
}

