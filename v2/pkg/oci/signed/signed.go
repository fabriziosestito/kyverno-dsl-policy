package signed

import _ "unsafe"

type Embedme interface{}

type image struct {Embedme}

func Image(i interface{}) interface{} {
 panic("stub")
}

type v1Index v1.ImageIndex

type index struct {Embedme}

func ImageIndex(i interface{}) interface{} {
 panic("stub")
}

func (ii *index) SignedImage(h interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (ii *index) SignedImageIndex(h interface{}) (interface{}, interface{}) {
 panic("stub")
}

