package publickey

import _ "unsafe"

type Embedme interface{}

type NamedWriter struct {Name interface{}; Embedme}

type Pkopts struct {KeyRef interface{}; Sk interface{}; Slot interface{}}

func GetPublicKey(ctx interface{}, opts interface{}, writer interface{}, pf interface{}) interface{} {
 panic("stub")
}

