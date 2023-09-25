package kubernetes

import _ "unsafe"

type Embedme interface{}

func client() (interface{}, interface{}) {
 panic("stub")
}

func checkImmutableSecretSupported(client interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetKeyPairSecret(ctx interface{}, k8sRef interface{}) (interface{}, interface{}) {
 panic("stub")
}

func KeyPairSecret(ctx interface{}, k8sRef interface{}, pf interface{}) interface{} {
 panic("stub")
}

func secret(keys interface{}, namespace, name interface{}, data interface{}, immutable interface{}) interface{} {
 panic("stub")
}

func parseRef(k8sRef interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

