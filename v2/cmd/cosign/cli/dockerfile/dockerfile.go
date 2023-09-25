package dockerfile

import _ "unsafe"

type Embedme interface{}

type VerifyDockerfileCommand struct {Embedme; BaseOnly interface{}}

func (c *VerifyDockerfileCommand) Exec(ctx interface{}, args interface{}) interface{} {
 panic("stub")
}

func getImagesFromDockerfile(ctx interface{}, dockerfile interface{}) (interface{}, interface{}) {
 panic("stub")
}

func getImageFromLine(line interface{}) interface{} {
 panic("stub")
}

