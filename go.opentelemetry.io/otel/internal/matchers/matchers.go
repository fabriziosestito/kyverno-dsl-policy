package matchers

import _ "unsafe"

type Embedme interface{}

type Expectation struct {t interface{}; actual interface{}}

func (e *Expectation) ToEqual(expected interface{}) {
 panic("stub")
}

func (e *Expectation) NotToEqual(expected interface{}) {
 panic("stub")
}

func (e *Expectation) ToBeNil() {
 panic("stub")
}

func (e *Expectation) NotToBeNil() {
 panic("stub")
}

func (e *Expectation) ToBeTrue() {
 panic("stub")
}

func (e *Expectation) ToBeFalse() {
 panic("stub")
}

func (e *Expectation) NotToPanic() {
 panic("stub")
}

func (e *Expectation) ToSucceed() {
 panic("stub")
}

func (e *Expectation) ToMatchError(expected interface{}) {
 panic("stub")
}

func (e *Expectation) ToContain(expected interface{}) {
 panic("stub")
}

func (e *Expectation) NotToContain(expected interface{}) {
 panic("stub")
}

func (e *Expectation) ToMatchInAnyOrder(expected interface{}) {
 panic("stub")
}

func (e *Expectation) ToBeTemporally(matcher interface{}, compareTo interface{}) {
 panic("stub")
}

func (e *Expectation) verifyExpectedNotNil(expected interface{}) {
 panic("stub")
}

func (e *Expectation) fail(msg interface{}) {
 panic("stub")
}

type Expecter struct {t interface{}}

func NewExpecter(t interface{}) interface{} {
 panic("stub")
}

func (a *Expecter) Expect(actual interface{}) interface{} {
 panic("stub")
}

type TemporalMatcher byte

