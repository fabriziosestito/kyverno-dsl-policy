package envvar

import _ "unsafe"

type Embedme interface{}

type envvar struct {}

func init() {
 panic("stub")
}

func (p *envvar) Enabled(interface{}) interface{} {
 panic("stub")
}

func (p *envvar) Provide(interface{}, interface{}) (interface{}, interface{}) {
 panic("stub")
}

