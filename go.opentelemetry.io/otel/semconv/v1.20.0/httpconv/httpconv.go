package httpconv

import _ "unsafe"

type Embedme interface{}

func ClientResponse(resp interface{}) interface{} {
 panic("stub")
}

func ClientRequest(req interface{}) interface{} {
 panic("stub")
}

func ClientStatus(code interface{}) (interface{}, interface{}) {
 panic("stub")
}

func ServerRequest(server interface{}, req interface{}) interface{} {
 panic("stub")
}

func ServerStatus(code interface{}) (interface{}, interface{}) {
 panic("stub")
}

func RequestHeader(h interface{}) interface{} {
 panic("stub")
}

func ResponseHeader(h interface{}) interface{} {
 panic("stub")
}

