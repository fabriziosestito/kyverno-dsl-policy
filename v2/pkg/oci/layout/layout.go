package layout

import _ "unsafe"

type Embedme interface{}

type v1Index v1.ImageIndex

type index struct {Embedme}

func SignedImageIndex(path interface{}) (interface{}, interface{}) {
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

func (i *index) imageByAnnotation(annotation interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (i *index) imageIndexByAnnotation(annotation interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (i *index) SignedImageIndex(h interface{}) (interface{}, interface{}) {
 panic("stub")
}

type sigs struct {Embedme}

func (s *sigs) Get() (interface{}, interface{}) {
 panic("stub")
}

func WriteSignedImage(path interface{}, si interface{}) interface{} {
 panic("stub")
}

func WriteSignedImageIndex(path interface{}, si interface{}) interface{} {
 panic("stub")
}

func writeSignedEntity(path interface{}, se interface{}) interface{} {
 panic("stub")
}

func isEmpty(s interface{}) interface{} {
 panic("stub")
}

func appendImage(path interface{}, img interface{}, annotation interface{}) interface{} {
 panic("stub")
}

