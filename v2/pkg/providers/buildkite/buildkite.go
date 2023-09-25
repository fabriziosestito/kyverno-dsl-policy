package buildkite

import _ "unsafe"

type Embedme interface{}

type buildkiteAgent struct {}

func init() {
 panic("stub")
}

func (ba *buildkiteAgent) Enabled(_ interface{}) interface{} {
 panic("stub")
}

func (ba *buildkiteAgent) Provide(ctx interface{}, audience interface{}) (interface{}, interface{}) {
 panic("stub")
}

