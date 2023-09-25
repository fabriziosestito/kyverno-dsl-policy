package term

import _ "unsafe"

type Embedme interface{}

type wordWrapWriter struct {limit interface{}; writer interface{}}

type TerminalSize struct {Width interface{}; Height interface{}}

func NewResponsiveWriter(w interface{}) interface{} {
 panic("stub")
}

func NewWordWrapWriter(w interface{}, limit interface{}) interface{} {
 panic("stub")
}

func getTerminalLimitWidth(terminalSize interface{}) interface{} {
 panic("stub")
}

func GetSize(fd interface{}) interface{} {
 panic("stub")
}

func GetWordWrapperLimit() (interface{}, interface{}) {
 panic("stub")
}

func (w wordWrapWriter) Write(p interface{}) (nn interface{}, err interface{}) {
 panic("stub")
}

