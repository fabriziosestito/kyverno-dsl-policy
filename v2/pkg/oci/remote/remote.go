package remote

import _ "unsafe"

type Embedme interface{}

func ResolveDigest(ref interface{}, opts interface{}) (interface{}, interface{}) {
	panic("stub")
}

type image struct {
	Embedme
	opt interface{}
}

func SignedImage(ref interface{}, options interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (i *image) Signatures() (interface{}, interface{}) {
	panic("stub")
}

func (i *image) Attestations() (interface{}, interface{}) {
	panic("stub")
}

func (i *image) Attachment(name interface{}) (interface{}, interface{}) {
	panic("stub")
}

type v1Index struct{}

type index struct {
	Embedme
	ref interface{}
	opt interface{}
}

func SignedImageIndex(ref interface{}, options interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (i *index) Signatures() (interface{}, interface{}) {
	panic("stub")
}

func (i *index) Attestations() (interface{}, interface{}) {
	panic("stub")
}

func (i *index) Attachment(name interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (i *index) SignedImage(h interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (i *index) SignedImageIndex(h interface{}) (interface{}, interface{}) {
	panic("stub")
}

type options struct {
	SignatureSuffix   interface{}
	AttestationSuffix interface{}
	SBOMSuffix        interface{}
	TagPrefix         interface{}
	TargetRepository  interface{}
	ROpt              interface{}
	NameOpts          interface{}
	OriginalOptions   interface{}
}

type Option func(interface{})

func makeOptions(target interface{}, opts interface{}) interface{} {
	panic("stub")
}

func WithPrefix(prefix interface{}) interface{} {
	panic("stub")
}

func WithSignatureSuffix(suffix interface{}) interface{} {
	panic("stub")
}

func WithAttestationSuffix(suffix interface{}) interface{} {
	panic("stub")
}

func WithSBOMSuffix(suffix interface{}) interface{} {
	panic("stub")
}

func WithRemoteOptions(opts interface{}) interface{} {
	panic("stub")
}

func WithTargetRepository(repo interface{}) interface{} {
	panic("stub")
}

func GetEnvTargetRepository() (interface{}, interface{}) {
	panic("stub")
}

func WithNameOptions(opts interface{}) interface{} {
	panic("stub")
}

func Referrers(d interface{}, artifactType interface{}, opts interface{}) (interface{}, interface{}) {
	panic("stub")
}

type EntityNotFoundError struct{ baseErr interface{} }

type attached struct {
	Embedme
	layer interface{}
}

func (e *EntityNotFoundError) Error() interface{} {
	panic("stub")
}

func NewEntityNotFoundError(err interface{}) interface{} {
	panic("stub")
}

func SignedEntity(ref interface{}, options interface{}) (interface{}, interface{}) {
	panic("stub")
}

func normalize(h interface{}, prefix interface{}, suffix interface{}) interface{} {
	panic("stub")
}

func SignatureTag(ref interface{}, opts interface{}) (interface{}, interface{}) {
	panic("stub")
}

func AttestationTag(ref interface{}, opts interface{}) (interface{}, interface{}) {
	panic("stub")
}

func SBOMTag(ref interface{}, opts interface{}) (interface{}, interface{}) {
	panic("stub")
}

func suffixTag(ref interface{}, suffix interface{}, o interface{}) (interface{}, interface{}) {
	panic("stub")
}

func signatures(digestable interface{}, o interface{}) (interface{}, interface{}) {
	panic("stub")
}

func attestations(digestable interface{}, o interface{}) (interface{}, interface{}) {
	panic("stub")
}

func attachment(digestable interface{}, attName interface{}, o interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (f *attached) FileMediaType() (interface{}, interface{}) {
	panic("stub")
}

func (f *attached) Payload() (interface{}, interface{}) {
	panic("stub")
}

func attachmentExperimentalOCI(digestable interface{}, attName interface{}, o interface{}) (interface{}, interface{}) {
	panic("stub")
}

type sigs struct{ Embedme }

func Signatures(ref interface{}, opts interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (s *sigs) Get() (interface{}, interface{}) {
	panic("stub")
}

type unknown struct {
	digest interface{}
	opt    interface{}
}

func SignedUnknown(digest interface{}, options interface{}) interface{} {
	panic("stub")
}

func (i *unknown) Digest() (interface{}, interface{}) {
	panic("stub")
}

func (i *unknown) Signatures() (interface{}, interface{}) {
	panic("stub")
}

func (i *unknown) Attestations() (interface{}, interface{}) {
	panic("stub")
}

func (i *unknown) Attachment(name interface{}) (interface{}, interface{}) {
	panic("stub")
}

type taggableManifest struct {
	raw       interface{}
	mediaType interface{}
}

func WriteSignedImageIndexImages(ref interface{}, sii interface{}, opts interface{}) interface{} {
	panic("stub")
}

func WriteSignatures(repo interface{}, se interface{}, opts interface{}) interface{} {
	panic("stub")
}

func WriteAttestations(repo interface{}, se interface{}, opts interface{}) interface{} {
	panic("stub")
}

func WriteSignaturesExperimentalOCI(d interface{}, se interface{}, opts interface{}) interface{} {
	panic("stub")
}

func (taggable taggableManifest) RawManifest() (interface{}, interface{}) {
	panic("stub")
}

func (taggable taggableManifest) MediaType() (interface{}, interface{}) {
	panic("stub")
}
