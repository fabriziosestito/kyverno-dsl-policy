package internal

import _ "unsafe"

type Embedme interface{}

type codeRange struct {fromInclusive interface{}; toInclusive interface{}}

type SemanticConventions struct {EnduserIDKey interface{}; HTTPClientIPKey interface{}; HTTPFlavorKey interface{}; HTTPHostKey interface{}; HTTPMethodKey interface{}; HTTPRequestContentLengthKey interface{}; HTTPRouteKey interface{}; HTTPSchemeHTTP interface{}; HTTPSchemeHTTPS interface{}; HTTPServerNameKey interface{}; HTTPStatusCodeKey interface{}; HTTPTargetKey interface{}; HTTPURLKey interface{}; HTTPUserAgentKey interface{}; NetHostIPKey interface{}; NetHostNameKey interface{}; NetHostPortKey interface{}; NetPeerIPKey interface{}; NetPeerNameKey interface{}; NetPeerPortKey interface{}; NetTransportIP interface{}; NetTransportOther interface{}; NetTransportTCP interface{}; NetTransportUDP interface{}; NetTransportUnix interface{}}

func (sc *SemanticConventions) NetAttributesFromHTTPRequest(network interface{}, request interface{}) interface{} {
 panic("stub")
}

func hostIPNamePort(hostWithPort interface{}) (ip interface{}, name interface{}, port interface{}) {
 panic("stub")
}

func (sc *SemanticConventions) EndUserAttributesFromHTTPRequest(request interface{}) interface{} {
 panic("stub")
}

func (sc *SemanticConventions) HTTPClientAttributesFromHTTPRequest(request interface{}) interface{} {
 panic("stub")
}

func (sc *SemanticConventions) httpCommonAttributesFromHTTPRequest(request interface{}) interface{} {
 panic("stub")
}

func (sc *SemanticConventions) httpBasicAttributesFromHTTPRequest(request interface{}) interface{} {
 panic("stub")
}

func (sc *SemanticConventions) HTTPServerMetricAttributesFromHTTPRequest(serverName interface{}, request interface{}) interface{} {
 panic("stub")
}

func (sc *SemanticConventions) HTTPServerAttributesFromHTTPRequest(serverName, route interface{}, request interface{}) interface{} {
 panic("stub")
}

func (sc *SemanticConventions) HTTPAttributesFromHTTPStatusCode(code interface{}) interface{} {
 panic("stub")
}

func (r codeRange) contains(code interface{}) interface{} {
 panic("stub")
}

func SpanStatusFromHTTPStatusCode(code interface{}) (interface{}, interface{}) {
 panic("stub")
}

func SpanStatusFromHTTPStatusCodeAndSpanKind(code interface{}, spanKind interface{}) (interface{}, interface{}) {
 panic("stub")
}

func validateHTTPStatusCode(code interface{}) (interface{}, interface{}) {
 panic("stub")
}

