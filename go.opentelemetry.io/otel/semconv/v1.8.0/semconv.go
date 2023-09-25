package semconv

import _ "unsafe"

type Embedme interface{}

func NetAttributesFromHTTPRequest(network interface{}, request interface{}) interface{} {
 panic("stub")
}

func EndUserAttributesFromHTTPRequest(request interface{}) interface{} {
 panic("stub")
}

func HTTPClientAttributesFromHTTPRequest(request interface{}) interface{} {
 panic("stub")
}

func HTTPServerMetricAttributesFromHTTPRequest(serverName interface{}, request interface{}) interface{} {
 panic("stub")
}

func HTTPServerAttributesFromHTTPRequest(serverName, route interface{}, request interface{}) interface{} {
 panic("stub")
}

func HTTPAttributesFromHTTPStatusCode(code interface{}) interface{} {
 panic("stub")
}

func SpanStatusFromHTTPStatusCode(code interface{}) (interface{}, interface{}) {
 panic("stub")
}

func SpanStatusFromHTTPStatusCodeAndSpanKind(code interface{}, spanKind interface{}) (interface{}, interface{}) {
 panic("stub")
}

