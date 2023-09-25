package internaltest

import _ "unsafe"

type Embedme interface{}

type FieldOffset struct {Name interface{}; Offset interface{}}

func Aligned8Byte(fields interface{}, out interface{}) interface{} {
 panic("stub")
}

type envStore struct {store interface{}}

type Env struct {Name interface{}; Value interface{}; Exists interface{}}

type EnvStore interface {}

func (s *envStore) add(env interface{}) {
 panic("stub")
}

func (s *envStore) Restore() interface{} {
 panic("stub")
}

func (s *envStore) setEnv(key, value interface{}) interface{} {
 panic("stub")
}

func (s *envStore) Record(key interface{}) {
 panic("stub")
}

func NewEnvStore() interface{} {
 panic("stub")
}

func newEnvStore() interface{} {
 panic("stub")
}

func SetEnvVariables(env interface{}) (interface{}, interface{}) {
 panic("stub")
}

type TestError string

func NewTestError(s interface{}) interface{} {
 panic("stub")
}

func (e TestError) Error() interface{} {
 panic("stub")
}

type Harness struct {t interface{}}

type testCtxKey struct {}

func NewHarness(t interface{}) interface{} {
 panic("stub")
}

func (h *Harness) TestTracerProvider(subjectFactory interface{}) {
 panic("stub")
}

func (h *Harness) TestTracer(subjectFactory interface{}) {
 panic("stub")
}

func (h *Harness) testSpan(tracerFactory interface{}) {
 panic("stub")
}

type TextMapCarrier struct {mtx interface{}; gets interface{}; sets interface{}; data interface{}}

func NewTextMapCarrier(data interface{}) interface{} {
 panic("stub")
}

func (c *TextMapCarrier) Keys() interface{} {
 panic("stub")
}

func (c *TextMapCarrier) Get(key interface{}) interface{} {
 panic("stub")
}

func (c *TextMapCarrier) GotKey(t interface{}, key interface{}) interface{} {
 panic("stub")
}

func (c *TextMapCarrier) GotN(t interface{}, n interface{}) interface{} {
 panic("stub")
}

func (c *TextMapCarrier) Set(key, value interface{}) {
 panic("stub")
}

func (c *TextMapCarrier) SetKeyValue(t interface{}, key, value interface{}) interface{} {
 panic("stub")
}

func (c *TextMapCarrier) SetN(t interface{}, n interface{}) interface{} {
 panic("stub")
}

func (c *TextMapCarrier) Reset(data interface{}) {
 panic("stub")
}

type state struct {Injections interface{}; Extractions interface{}}

type TextMapPropagator struct {name interface{}; ctxKey interface{}}

type ctxKeyType string

func newState(encoded interface{}) interface{} {
 panic("stub")
}

func (s state) String() interface{} {
 panic("stub")
}

func NewTextMapPropagator(name interface{}) interface{} {
 panic("stub")
}

func (p *TextMapPropagator) stateFromContext(ctx interface{}) interface{} {
 panic("stub")
}

func (p *TextMapPropagator) stateFromCarrier(carrier interface{}) interface{} {
 panic("stub")
}

func (p *TextMapPropagator) Inject(ctx interface{}, carrier interface{}) {
 panic("stub")
}

func (p *TextMapPropagator) InjectedN(t interface{}, carrier interface{}, n interface{}) interface{} {
 panic("stub")
}

func (p *TextMapPropagator) Extract(ctx interface{}, carrier interface{}) interface{} {
 panic("stub")
}

func (p *TextMapPropagator) ExtractedN(t interface{}, ctx interface{}, n interface{}) interface{} {
 panic("stub")
}

func (p *TextMapPropagator) Fields() interface{} {
 panic("stub")
}

