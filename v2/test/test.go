package test

import _ "unsafe"

type Embedme interface{}

func createCertificate(template interface{}, parent interface{}, pub interface{}, priv interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GenerateRootCa() (interface{}, interface{}, interface{}) {
 panic("stub")
}

func GenerateSubordinateCa(rootTemplate interface{}, rootPriv interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func GenerateLeafCertWithExpiration(subject interface{}, oidcIssuer interface{}, expiration interface{}, priv interface{}, parentTemplate interface{}, parentPriv interface{}) (interface{}, interface{}) {
 panic("stub")
}

func GenerateLeafCert(subject interface{}, oidcIssuer interface{}, parentTemplate interface{}, parentPriv interface{}, exts interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func GenerateLeafCertWithGitHubOIDs(subject interface{}, oidcIssuer interface{}, githubWorkflowTrigger, githubWorkflowSha, githubWorkflowName, githubWorkflowRepository, githubWorkflowRef interface{}, parentTemplate interface{}, parentPriv interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

func GenerateLeafCertWithSubjectAlternateNames(dnsNames interface{}, emailAddresses interface{}, ipAddresses interface{}, uris interface{}, oidcIssuer interface{}, parentTemplate interface{}, parentPriv interface{}) (interface{}, interface{}, interface{}) {
 panic("stub")
}

