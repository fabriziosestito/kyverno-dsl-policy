package internal

import _ "unsafe"

type Embedme interface{}

type codeRange struct {fromInclusive interface{}; toInclusive interface{}}

type HTTPConv struct {NetConv interface{}; EnduserIDKey interface{}; HTTPClientIPKey interface{}; HTTPFlavorKey interface{}; HTTPMethodKey interface{}; HTTPRequestContentLengthKey interface{}; HTTPResponseContentLengthKey interface{}; HTTPRouteKey interface{}; HTTPSchemeHTTP interface{}; HTTPSchemeHTTPS interface{}; HTTPStatusCodeKey interface{}; HTTPTargetKey interface{}; HTTPURLKey interface{}; HTTPUserAgentKey interface{}}

func (c *HTTPConv) ClientResponse(resp interface{}) interface{} {
 panic("stub")
}

func (c *HTTPConv) ClientRequest(req interface{}) interface{} {
 panic("stub")
}

func (c *HTTPConv) ServerRequest(server interface{}, req interface{}) interface{} {
 panic("stub")
}

func (c *HTTPConv) method(method interface{}) interface{} {
 panic("stub")
}

func (c *HTTPConv) scheme(https interface{}) interface{} {
 panic("stub")
}

func (c *HTTPConv) proto(proto interface{}) interface{} {
 panic("stub")
}

func serverClientIP(xForwardedFor interface{}) interface{} {
 panic("stub")
}

func requiredHTTPPort(https interface{}, port interface{}) interface{} {
 panic("stub")
}

func firstHostPort(source interface{}) (host interface{}, port interface{}) {
 panic("stub")
}

func (c *HTTPConv) RequestHeader(h interface{}) interface{} {
 panic("stub")
}

func (c *HTTPConv) ResponseHeader(h interface{}) interface{} {
 panic("stub")
}

func (c *HTTPConv) header(prefix interface{}, h interface{}) interface{} {
 panic("stub")
}

func (c *HTTPConv) ClientStatus(code interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (c *HTTPConv) ServerStatus(code interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (r codeRange) contains(code interface{}) interface{} {
 panic("stub")
}

func validateHTTPStatusCode(code interface{}) (interface{}, interface{}) {
 panic("stub")
}

type NetConv struct {NetHostNameKey interface{}; NetHostPortKey interface{}; NetPeerNameKey interface{}; NetPeerPortKey interface{}; NetSockFamilyKey interface{}; NetSockPeerAddrKey interface{}; NetSockPeerPortKey interface{}; NetSockHostAddrKey interface{}; NetSockHostPortKey interface{}; NetTransportOther interface{}; NetTransportTCP interface{}; NetTransportUDP interface{}; NetTransportInProc interface{}}

func (c *NetConv) Transport(network interface{}) interface{} {
 panic("stub")
}

func (c *NetConv) Host(address interface{}) interface{} {
 panic("stub")
}

func (c *NetConv) Server(address interface{}, ln interface{}) interface{} {
 panic("stub")
}

func (c *NetConv) HostName(name interface{}) interface{} {
 panic("stub")
}

func (c *NetConv) HostPort(port interface{}) interface{} {
 panic("stub")
}

func (c *NetConv) Client(address interface{}, conn interface{}) interface{} {
 panic("stub")
}

func family(network, address interface{}) interface{} {
 panic("stub")
}

func nonZeroStr(strs interface{}) interface{} {
 panic("stub")
}

func positiveInt(ints interface{}) interface{} {
 panic("stub")
}

func (c *NetConv) Peer(address interface{}) interface{} {
 panic("stub")
}

func (c *NetConv) PeerName(name interface{}) interface{} {
 panic("stub")
}

func (c *NetConv) PeerPort(port interface{}) interface{} {
 panic("stub")
}

func (c *NetConv) SockPeerAddr(addr interface{}) interface{} {
 panic("stub")
}

func (c *NetConv) SockPeerPort(port interface{}) interface{} {
 panic("stub")
}

func splitHostPort(hostport interface{}) (host interface{}, port interface{}) {
 panic("stub")
}

