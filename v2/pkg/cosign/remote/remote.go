package remote

import _ "unsafe"

type Embedme interface{}

type File interface {}

type file struct {path interface{}; platform interface{}}

type MediaTypeGetter func(b interface{}) interface{}

func FilesFromFlagList(sl interface{}) interface{} {
 panic("stub")
}

func FileFromFlag(s interface{}) interface{} {
 panic("stub")
}

func (f *file) Path() interface{} {
 panic("stub")
}

func (f *file) Contents() (interface{}, interface{}) {
 panic("stub")
}

func (f *file) Platform() interface{} {
 panic("stub")
}

func (f *file) String() interface{} {
 panic("stub")
}

func DefaultMediaTypeGetter(b interface{}) interface{} {
 panic("stub")
}

func UploadFiles(ref interface{}, files interface{}, annotations interface{}, getMt interface{}, remoteOpts interface{}) (interface{}, interface{}) {
 panic("stub")
}

type dd struct {verifier interface{}}

type ro struct {predicateURI interface{}}

type replaceOCISignatures struct {Embedme; attestations interface{}}

func NewDupeDetector(v interface{}) interface{} {
 panic("stub")
}

func NewReplaceOp(predicateURI interface{}) interface{} {
 panic("stub")
}

func (dd *dd) Find(sigImage interface{}, newSig interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (r *ro) Replace(signatures interface{}, o interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (r *replaceOCISignatures) Get() (interface{}, interface{}) {
 panic("stub")
}

