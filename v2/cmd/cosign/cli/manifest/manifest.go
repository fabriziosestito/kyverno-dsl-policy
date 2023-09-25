package manifest

import _ "unsafe"

type Embedme interface{}

type VerifyManifestCommand struct {Embedme}

type unionImagesKind struct {Spec interface{}}

type imageContainers struct {Containers interface{}; InitContainers interface{}}

func (c *VerifyManifestCommand) Exec(ctx interface{}, args interface{}) interface{} {
 panic("stub")
}

func (uik *unionImagesKind) images() interface{} {
 panic("stub")
}

func getImagesFromYamlManifest(manifest interface{}) (interface{}, interface{}) {
 panic("stub")
}

func isExtensionAllowed(ext interface{}) interface{} {
 panic("stub")
}

func allowedExtensionsForManifest() interface{} {
 panic("stub")
}

