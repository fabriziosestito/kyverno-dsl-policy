package ui

import _ "unsafe"

type Embedme interface{}

type callbackFunc func(interface{}, interface{})

type ctxKey struct {}

type WriteFunc func(interface{})

type Env struct {Stderr interface{}; Stdin interface{}}

func defaultEnv() interface{} {
 panic("stub")
}

func (c ctxKey) String() interface{} {
 panic("stub")
}

func getEnv(ctx interface{}) interface{} {
 panic("stub")
}

func WithEnv(ctx interface{}, e interface{}) interface{} {
 panic("stub")
}

func RunWithTestCtx(callback interface{}) interface{} {
 panic("stub")
}

func (w *Env) infof(msg interface{}, a interface{}) {
 panic("stub")
}

func Infof(ctx interface{}, msg interface{}, a interface{}) {
 panic("stub")
}

func (w *Env) warnf(msg interface{}, a interface{}) {
 panic("stub")
}

func Warnf(ctx interface{}, msg interface{}, a interface{}) {
 panic("stub")
}

type ErrPromptDeclined struct {}

type ErrInvalidInput struct {Got interface{}; Allowed interface{}}

func (e *ErrPromptDeclined) Error() interface{} {
 panic("stub")
}

func (e *ErrInvalidInput) Error() interface{} {
 panic("stub")
}

func newInvalidYesOrNoInput(got interface{}) interface{} {
 panic("stub")
}

func (w *Env) prompt() interface{} {
 panic("stub")
}

func ConfirmContinue(ctx interface{}) interface{} {
 panic("stub")
}

