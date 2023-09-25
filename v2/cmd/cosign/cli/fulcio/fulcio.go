package fulcio

import _ "unsafe"

type Embedme interface{}

type Signer struct {Cert interface{}; Chain interface{}; SCT interface{}; Embedme}

type oidcConnector interface {}

type realConnector struct {flow interface{}}

func (rf *realConnector) OIDConnect(url, clientID, secret, redirectURL interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getCertForOauthID(sv interface{}, fc interface{}, connector interface{}, oidcIssuer, oidcClientID, oidcClientSecret, oidcRedirectURL interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetCert(_ interface{}, sv interface{}, idToken, flow, oidcIssuer, oidcClientID, oidcClientSecret, oidcRedirectURL interface{}, fClient interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewSigner(ctx interface{}, ko interface{}, signer interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (f *Signer) PublicKey(opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GetRoots() (interface{}, interface{}) {
 panic("stub")
}

func GetIntermediates() (interface{}, interface{}) {
 panic("stub")
}

func NewClient(fulcioURL interface{}) (interface{}, interface{}) {
 panic("stub")
}

func idToken(s interface{}) (interface{}, interface{}) {
 panic("stub")
}

