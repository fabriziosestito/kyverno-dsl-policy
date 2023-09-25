package mutate

import _ "unsafe"

type Embedme interface{}

type mapPassKey struct {}

type Fn func(interface{}, interface{}) (interface{}, interface{})

func Map(ctx interface{}, parent interface{}, fn interface{}) (interface{}, interface{}) {
 panic("stub")
}

func before(ctx interface{}) interface{} {
 panic("stub")
}

func after(ctx interface{}) interface{} {
 panic("stub")
}

func IsBeforeChildren(ctx interface{}) interface{} {
 panic("stub")
}

func IsAfterChildren(ctx interface{}) interface{} {
 panic("stub")
}

type IndexAddendum struct {Add interface{}; Embedme}

type indexWrapper struct {Embedme; ogbase interface{}; addendum interface{}}

type signedImageIndex struct {Embedme; sig interface{}; att interface{}; so interface{}; attachments interface{}}

type Appendable interface {}

type signedImage struct {Embedme; sig interface{}; att interface{}; so interface{}; attachments interface{}}

type signedUnknown struct {Embedme; sig interface{}; att interface{}; so interface{}; attachments interface{}}

type digestable interface {}

type v1Index v1.ImageIndex

type ociSignedImageIndex oci.SignedImageIndex

func AppendManifests(base interface{}, adds interface{}) interface{} {
 panic("stub")
}

func (i *indexWrapper) Signatures() (interface{}, interface{}) {
 panic("stub")
}

func (i *indexWrapper) Attestations() (interface{}, interface{}) {
 panic("stub")
}

func (i *indexWrapper) SignedImage(h interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (i *indexWrapper) SignedImageIndex(h interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachSignatureToEntity(se interface{}, sig interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachAttestationToEntity(se interface{}, att interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachFileToEntity(se interface{}, name interface{}, f interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachSignatureToImage(si interface{}, sig interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachAttestationToImage(si interface{}, att interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachFileToImage(si interface{}, name interface{}, f interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (si *signedImage) Signatures() (interface{}, interface{}) {
 panic("stub")
}

func (si *signedImage) Attestations() (interface{}, interface{}) {
 panic("stub")
}

func (si *signedImage) Attachment(attName interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachSignatureToImageIndex(sii interface{}, sig interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachAttestationToImageIndex(sii interface{}, att interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachFileToImageIndex(sii interface{}, name interface{}, f interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (sii *signedImageIndex) Signatures() (interface{}, interface{}) {
 panic("stub")
}

func (sii *signedImageIndex) Attestations() (interface{}, interface{}) {
 panic("stub")
}

func (sii *signedImageIndex) Attachment(attName interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachSignatureToUnknown(se interface{}, sig interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachAttestationToUnknown(se interface{}, att interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func AttachFileToUnknown(se interface{}, name interface{}, f interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (si *signedUnknown) Digest() (interface{}, interface{}) {
 panic("stub")
}

func (si *signedUnknown) Signatures() (interface{}, interface{}) {
 panic("stub")
}

func (si *signedUnknown) Attestations() (interface{}, interface{}) {
 panic("stub")
}

func (si *signedUnknown) Attachment(attName interface{}) (interface{}, interface{}) {
 panic("stub")
}

type ReplaceOp interface {}

type signOpts struct {dd interface{}; ro interface{}}

type SignatureOption func(interface{})

type SignOption func(interface{})

type signatureOpts struct {annotations interface{}; bundle interface{}; rfc3161Timestamp interface{}; cert interface{}; chain interface{}; mediaType interface{}}

type DupeDetector interface {}

func makeSignOpts(opts interface{}) interface{} {
 panic("stub")
}

func WithDupeDetector(dd interface{}) interface{} {
 panic("stub")
}

func WithReplaceOp(ro interface{}) interface{} {
 panic("stub")
}

func WithAnnotations(annotations interface{}) interface{} {
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

func WithMediaType(mediaType interface{}) interface{} {
 panic("stub")
}

func makeSignatureOption(opts interface{}) interface{} {
 panic("stub")
}

type sigWrapper struct {wrapped interface{}; annotations interface{}; bundle interface{}; rfc3161Timestamp interface{}; cert interface{}; chain interface{}; mediaType interface{}}

func copyAnnotations(ann interface{}) interface{} {
 panic("stub")
}

func (sw *sigWrapper) Annotations() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) Payload() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) Signature() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) Base64Signature() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) Cert() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) Chain() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) Bundle() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) RFC3161Timestamp() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) MediaType() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) Digest() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) DiffID() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) Compressed() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) Uncompressed() (interface{}, interface{}) {
 panic("stub")
}

func (sw *sigWrapper) Size() (interface{}, interface{}) {
 panic("stub")
}

func Signature(original interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

type sigAppender struct {Embedme; base interface{}; sigs interface{}}

func AppendSignatures(base interface{}, sigs interface{}) (interface{}, interface{}) {
 panic("stub")
}

func ReplaceSignatures(base interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (sa *sigAppender) Get() (interface{}, interface{}) {
 panic("stub")
}

