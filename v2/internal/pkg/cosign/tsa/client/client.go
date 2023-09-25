package client

import _ "unsafe"

type Embedme interface{}

type TimestampAuthorityClient interface {}

type TimestampAuthorityClientImpl struct {Embedme; URL interface{}; CACert interface{}; Cert interface{}; Key interface{}; ServerName interface{}; Timeout interface{}}

func getHTTPTransport(cacertFilename, certFilename, keyFilename, serverName interface{}, timeout interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (t *TimestampAuthorityClientImpl) GetTimestampResponse(tsq interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewTSAClient(url interface{}) interface{} {
 panic("stub")
}

func NewTSAClientMTLS(url, cacert, cert, key, serverName interface{}) interface{} {
 panic("stub")
}

