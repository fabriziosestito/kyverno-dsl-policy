package trace

import (
	"time"
	_ "unsafe"

	"go.opentelemetry.io/otel/attribute"
)

type Embedme interface{}

type TracerConfig struct {
	instrumentationVersion interface{}
	schemaURL              interface{}
	attrs                  interface{}
}

type SpanStartOption interface{}

type EventConfig struct {
	attributes interface{}
	timestamp  interface{}
	stackTrace interface{}
}

type EventOption interface{}

type SpanStartEventOption interface{}

type spanOptionFunc func(interface{}) interface{}

type SpanEndOption interface{}

type SpanConfig struct {
	attributes interface{}
	timestamp  interface{}
	links      interface{}
	newRoot    interface{}
	spanKind   interface{}
	stackTrace interface{}
}

type SpanOption interface{}

type TracerOption interface{}

type SpanEndEventOption interface{}

type SpanEventOption interface{}

type tracerOptionFunc func(interface{}) interface{}

type attributeOption []attribute.KeyValue

type timestampOption time.Time

type stackTraceOption bool

func (t *TracerConfig) InstrumentationVersion() interface{} {
	panic("stub")
}

func (t *TracerConfig) InstrumentationAttributes() interface{} {
	panic("stub")
}

func (t *TracerConfig) SchemaURL() interface{} {
	panic("stub")
}

func NewTracerConfig(options interface{}) interface{} {
	panic("stub")
}

func (fn tracerOptionFunc) apply(cfg interface{}) interface{} {
	panic("stub")
}

func (cfg *SpanConfig) Attributes() interface{} {
	panic("stub")
}

func (cfg *SpanConfig) Timestamp() interface{} {
	panic("stub")
}

func (cfg *SpanConfig) StackTrace() interface{} {
	panic("stub")
}

func (cfg *SpanConfig) Links() interface{} {
	panic("stub")
}

func (cfg *SpanConfig) NewRoot() interface{} {
	panic("stub")
}

func (cfg *SpanConfig) SpanKind() interface{} {
	panic("stub")
}

func NewSpanStartConfig(options interface{}) interface{} {
	panic("stub")
}

func NewSpanEndConfig(options interface{}) interface{} {
	panic("stub")
}

func (fn spanOptionFunc) applySpanStart(cfg interface{}) interface{} {
	panic("stub")
}

func (cfg *EventConfig) Attributes() interface{} {
	panic("stub")
}

func (cfg *EventConfig) Timestamp() interface{} {
	panic("stub")
}

func (cfg *EventConfig) StackTrace() interface{} {
	panic("stub")
}

func NewEventConfig(options interface{}) interface{} {
	panic("stub")
}

func (o attributeOption) applySpan(c interface{}) interface{} {
	panic("stub")
}

func (o attributeOption) applySpanStart(c interface{}) interface{} {
	panic("stub")
}

func (o attributeOption) applyEvent(c interface{}) interface{} {
	panic("stub")
}

func WithAttributes(attributes interface{}) interface{} {
	panic("stub")
}

func (o timestampOption) applySpan(c interface{}) interface{} {
	panic("stub")
}

func (o timestampOption) applySpanStart(c interface{}) interface{} {
	panic("stub")
}

func (o timestampOption) applySpanEnd(c interface{}) interface{} {
	panic("stub")
}

func (o timestampOption) applyEvent(c interface{}) interface{} {
	panic("stub")
}

func WithTimestamp(t interface{}) interface{} {
	panic("stub")
}

func (o stackTraceOption) applyEvent(c interface{}) interface{} {
	panic("stub")
}

func (o stackTraceOption) applySpan(c interface{}) interface{} {
	panic("stub")
}

func (o stackTraceOption) applySpanEnd(c interface{}) interface{} {
	panic("stub")
}

func WithStackTrace(b interface{}) interface{} {
	panic("stub")
}

func WithLinks(links interface{}) interface{} {
	panic("stub")
}

func WithNewRoot() interface{} {
	panic("stub")
}

func WithSpanKind(kind interface{}) interface{} {
	panic("stub")
}

func WithInstrumentationVersion(version interface{}) interface{} {
	panic("stub")
}

func WithInstrumentationAttributes(attr interface{}) interface{} {
	panic("stub")
}

func WithSchemaURL(schemaURL interface{}) interface{} {
	panic("stub")
}

type traceContextKeyType int

func ContextWithSpan(parent interface{}, span interface{}) interface{} {
	panic("stub")
}

func ContextWithSpanContext(parent interface{}, sc interface{}) interface{} {
	panic("stub")
}

func ContextWithRemoteSpanContext(parent interface{}, rsc interface{}) interface{} {
	panic("stub")
}

func SpanFromContext(ctx interface{}) interface{} {
	panic("stub")
}

func SpanContextFromContext(ctx interface{}) interface{} {
	panic("stub")
}

type nonRecordingSpan struct {
	Embedme
	sc interface{}
}

func (s nonRecordingSpan) SpanContext() interface{} {
	panic("stub")
}

type noopSpan struct{}

type noopTracerProvider struct{}

type noopTracer struct{}

func NewNoopTracerProvider() interface{} {
	panic("stub")
}

func (p noopTracerProvider) Tracer(interface{}, interface{}) interface{} {
	panic("stub")
}

func (t noopTracer) Start(ctx interface{}, name interface{}, _ interface{}) (interface{}, interface{}) {
	panic("stub")
}

type TraceID [16]byte

type SpanContextConfig struct {
	TraceID    interface{}
	SpanID     interface{}
	TraceFlags interface{}
	TraceState interface{}
	Remote     interface{}
}

type Tracer interface{}

type TraceFlags byte

type SpanKind int

type errorConst string

type Span interface{}

type Link struct {
	SpanContext interface{}
	Attributes  interface{}
}

type SpanID [8]byte

type SpanContext struct {
	traceID    interface{}
	spanID     interface{}
	traceFlags interface{}
	traceState interface{}
	remote     interface{}
}

type TracerProvider interface{}

func (e errorConst) Error() interface{} {
	panic("stub")
}

func (t TraceID) IsValid() interface{} {
	panic("stub")
}

func (t TraceID) MarshalJSON() (interface{}, interface{}) {
	panic("stub")
}

func (t TraceID) String() interface{} {
	panic("stub")
}

func (s SpanID) IsValid() interface{} {
	panic("stub")
}

func (s SpanID) MarshalJSON() (interface{}, interface{}) {
	panic("stub")
}

func (s SpanID) String() interface{} {
	panic("stub")
}

func TraceIDFromHex(h interface{}) (interface{}, interface{}) {
	panic("stub")
}

func SpanIDFromHex(h interface{}) (interface{}, interface{}) {
	panic("stub")
}

func decodeHex(h interface{}, b interface{}) interface{} {
	panic("stub")
}

func (tf TraceFlags) IsSampled() interface{} {
	panic("stub")
}

func (tf TraceFlags) WithSampled(sampled interface{}) interface{} {
	panic("stub")
}

func (tf TraceFlags) MarshalJSON() (interface{}, interface{}) {
	panic("stub")
}

func (tf TraceFlags) String() interface{} {
	panic("stub")
}

func NewSpanContext(config interface{}) interface{} {
	panic("stub")
}

func (sc SpanContext) IsValid() interface{} {
	panic("stub")
}

func (sc SpanContext) IsRemote() interface{} {
	panic("stub")
}

func (sc SpanContext) WithRemote(remote interface{}) interface{} {
	panic("stub")
}

func (sc SpanContext) TraceID() interface{} {
	panic("stub")
}

func (sc SpanContext) HasTraceID() interface{} {
	panic("stub")
}

func (sc SpanContext) WithTraceID(traceID interface{}) interface{} {
	panic("stub")
}

func (sc SpanContext) SpanID() interface{} {
	panic("stub")
}

func (sc SpanContext) HasSpanID() interface{} {
	panic("stub")
}

func (sc SpanContext) WithSpanID(spanID interface{}) interface{} {
	panic("stub")
}

func (sc SpanContext) TraceFlags() interface{} {
	panic("stub")
}

func (sc SpanContext) IsSampled() interface{} {
	panic("stub")
}

func (sc SpanContext) WithTraceFlags(flags interface{}) interface{} {
	panic("stub")
}

func (sc SpanContext) TraceState() interface{} {
	panic("stub")
}

func (sc SpanContext) WithTraceState(state interface{}) interface{} {
	panic("stub")
}

func (sc SpanContext) Equal(other interface{}) interface{} {
	panic("stub")
}

func (sc SpanContext) MarshalJSON() (interface{}, interface{}) {
	panic("stub")
}

func LinkFromContext(ctx interface{}, attrs interface{}) interface{} {
	panic("stub")
}

func ValidateSpanKind(spanKind interface{}) interface{} {
	panic("stub")
}

func (sk SpanKind) String() interface{} {
	panic("stub")
}

type TraceState struct{ list interface{} }

type member struct {
	Key   interface{}
	Value interface{}
}

func newMember(key, value interface{}) (interface{}, interface{}) {
	panic("stub")
}

func parseMember(m interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (m member) String() interface{} {
	panic("stub")
}

func ParseTraceState(tracestate interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (ts TraceState) MarshalJSON() (interface{}, interface{}) {
	panic("stub")
}

func (ts TraceState) String() interface{} {
	panic("stub")
}

func (ts TraceState) Get(key interface{}) interface{} {
	panic("stub")
}

func (ts TraceState) Insert(key, value interface{}) (interface{}, interface{}) {
	panic("stub")
}

func (ts TraceState) Delete(key interface{}) interface{} {
	panic("stub")
}

func (ts TraceState) Len() interface{} {
	panic("stub")
}
