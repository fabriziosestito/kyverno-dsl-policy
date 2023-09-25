package static

import _ "unsafe"

type Embedme interface{}

type file struct {Embedme; layer interface{}}

func NewFile(payload interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (f *file) FileMediaType() (interface{}, interface{}) {
 panic("stub")
}

func (f *file) Payload() (interface{}, interface{}) {
 panic("stub")
}

type options struct {LayerMediaType interface{}; ConfigMediaType interface{}; Bundle interface{}; RFC3161Timestamp interface{}; Cert interface{}; Chain interface{}; Annotations interface{}}

type Option func(interface{})

func makeOptions(opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func WithLayerMediaType(mt interface{}) interface{} {
 panic("stub")
}

func WithConfigMediaType(mt interface{}) interface{} {
 panic("stub")
}

func WithAnnotations(ann interface{}) interface{} {
 panic("stub")
}

func WithBundle(b interface{}) interface{} {
 panic("stub")
}

func WithRFC3161Timestamp(b interface{}) interface{} {
 panic("stub")
}

func WithCertChain(cert, chain interface{}) interface{} {
 panic("stub")
}

type staticLayer struct {b interface{}; b64sig interface{}; opts interface{}}

func NewSignature(payload interface{}, b64sig interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewAttestation(payload interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func Copy(sig interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Annotations() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Payload() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Signature() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Base64Signature() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Cert() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Chain() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Bundle() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) RFC3161Timestamp() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Digest() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) DiffID() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Compressed() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Uncompressed() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) Size() (interface{}, interface{}) {
 panic("stub")
}

func (l *staticLayer) MediaType() (interface{}, interface{}) {
 panic("stub")
}

