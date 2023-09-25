package global

import _ "unsafe"

type Embedme interface{}

type ErrDelegator struct {delegate interface{}}

type ErrLogger struct {l interface{}}

type ErrorHandler interface {}

func (d *ErrDelegator) Handle(err interface{}) {
 panic("stub")
}

func (d *ErrDelegator) getDelegate() interface{} {
 panic("stub")
}

func (d *ErrDelegator) setDelegate(eh interface{}) {
 panic("stub")
}

func defaultErrorHandler() interface{} {
 panic("stub")
}

func (h *ErrLogger) Handle(err interface{}) {
 panic("stub")
}

func GetErrorHandler() interface{} {
 panic("stub")
}

func SetErrorHandler(h interface{}) {
 panic("stub")
}

func Handle(err interface{}) {
 panic("stub")
}

type siCounter struct {Embedme; name interface{}; opts interface{}; delegate interface{}}

type unwrapper interface {}

type afCounter struct {Embedme; Embedme; name interface{}; opts interface{}; delegate interface{}}

type afGauge struct {Embedme; Embedme; name interface{}; opts interface{}; delegate interface{}}

type sfCounter struct {Embedme; name interface{}; opts interface{}; delegate interface{}}

type afUpDownCounter struct {Embedme; Embedme; name interface{}; opts interface{}; delegate interface{}}

type aiGauge struct {Embedme; Embedme; name interface{}; opts interface{}; delegate interface{}}

type sfUpDownCounter struct {Embedme; name interface{}; opts interface{}; delegate interface{}}

type sfHistogram struct {Embedme; name interface{}; opts interface{}; delegate interface{}}

type siUpDownCounter struct {Embedme; name interface{}; opts interface{}; delegate interface{}}

type aiCounter struct {Embedme; Embedme; name interface{}; opts interface{}; delegate interface{}}

type aiUpDownCounter struct {Embedme; Embedme; name interface{}; opts interface{}; delegate interface{}}

type siHistogram struct {Embedme; name interface{}; opts interface{}; delegate interface{}}

func (i *afCounter) setDelegate(m interface{}) {
 panic("stub")
}

func (i *afCounter) Unwrap() interface{} {
 panic("stub")
}

func (i *afUpDownCounter) setDelegate(m interface{}) {
 panic("stub")
}

func (i *afUpDownCounter) Unwrap() interface{} {
 panic("stub")
}

func (i *afGauge) setDelegate(m interface{}) {
 panic("stub")
}

func (i *afGauge) Unwrap() interface{} {
 panic("stub")
}

func (i *aiCounter) setDelegate(m interface{}) {
 panic("stub")
}

func (i *aiCounter) Unwrap() interface{} {
 panic("stub")
}

func (i *aiUpDownCounter) setDelegate(m interface{}) {
 panic("stub")
}

func (i *aiUpDownCounter) Unwrap() interface{} {
 panic("stub")
}

func (i *aiGauge) setDelegate(m interface{}) {
 panic("stub")
}

func (i *aiGauge) Unwrap() interface{} {
 panic("stub")
}

func (i *sfCounter) setDelegate(m interface{}) {
 panic("stub")
}

func (i *sfCounter) Add(ctx interface{}, incr interface{}, opts interface{}) {
 panic("stub")
}

func (i *sfUpDownCounter) setDelegate(m interface{}) {
 panic("stub")
}

func (i *sfUpDownCounter) Add(ctx interface{}, incr interface{}, opts interface{}) {
 panic("stub")
}

func (i *sfHistogram) setDelegate(m interface{}) {
 panic("stub")
}

func (i *sfHistogram) Record(ctx interface{}, x interface{}, opts interface{}) {
 panic("stub")
}

func (i *siCounter) setDelegate(m interface{}) {
 panic("stub")
}

func (i *siCounter) Add(ctx interface{}, x interface{}, opts interface{}) {
 panic("stub")
}

func (i *siUpDownCounter) setDelegate(m interface{}) {
 panic("stub")
}

func (i *siUpDownCounter) Add(ctx interface{}, x interface{}, opts interface{}) {
 panic("stub")
}

func (i *siHistogram) setDelegate(m interface{}) {
 panic("stub")
}

func (i *siHistogram) Record(ctx interface{}, x interface{}, opts interface{}) {
 panic("stub")
}

func init() {
 panic("stub")
}

func SetLogger(l interface{}) {
 panic("stub")
}

func getLogger() interface{} {
 panic("stub")
}

func Info(msg interface{}, keysAndValues interface{}) {
 panic("stub")
}

func Error(err interface{}, msg interface{}, keysAndValues interface{}) {
 panic("stub")
}

func Debug(msg interface{}, keysAndValues interface{}) {
 panic("stub")
}

func Warn(msg interface{}, keysAndValues interface{}) {
 panic("stub")
}

type meterProvider struct {Embedme; mtx interface{}; meters interface{}; delegate interface{}}

type meter struct {Embedme; name interface{}; opts interface{}; mtx interface{}; instruments interface{}; registry interface{}; delegate interface{}}

type delegatedInstrument interface {}

type wrapped interface {}

type registration struct {Embedme; instruments interface{}; function interface{}; unreg interface{}; unregMu interface{}}

func (p *meterProvider) setDelegate(provider interface{}) {
 panic("stub")
}

func (p *meterProvider) Meter(name interface{}, opts interface{}) interface{} {
 panic("stub")
}

func (m *meter) setDelegate(provider interface{}) {
 panic("stub")
}

func (m *meter) Int64Counter(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Int64UpDownCounter(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Int64Histogram(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Int64ObservableCounter(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Int64ObservableUpDownCounter(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Int64ObservableGauge(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Float64Counter(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Float64UpDownCounter(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Float64Histogram(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Float64ObservableCounter(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Float64ObservableUpDownCounter(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) Float64ObservableGauge(name interface{}, options interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *meter) RegisterCallback(f interface{}, insts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func unwrapInstruments(instruments interface{}) interface{} {
 panic("stub")
}

func (c *registration) setDelegate(m interface{}) {
 panic("stub")
}

func (c *registration) Unregister() interface{} {
 panic("stub")
}

type textMapPropagator struct {mtx interface{}; once interface{}; delegate interface{}; noop interface{}}

func newTextMapPropagator() interface{} {
 panic("stub")
}

func (p *textMapPropagator) SetDelegate(delegate interface{}) {
 panic("stub")
}

func (p *textMapPropagator) effectiveDelegate() interface{} {
 panic("stub")
}

func (p *textMapPropagator) Inject(ctx interface{}, carrier interface{}) {
 panic("stub")
}

func (p *textMapPropagator) Extract(ctx interface{}, carrier interface{}) interface{} {
 panic("stub")
}

func (p *textMapPropagator) Fields() interface{} {
 panic("stub")
}

type propagatorsHolder struct {tm interface{}}

type tracerProviderHolder struct {tp interface{}}

type meterProviderHolder struct {mp interface{}}

func TracerProvider() interface{} {
 panic("stub")
}

func SetTracerProvider(tp interface{}) {
 panic("stub")
}

func TextMapPropagator() interface{} {
 panic("stub")
}

func SetTextMapPropagator(p interface{}) {
 panic("stub")
}

func MeterProvider() interface{} {
 panic("stub")
}

func SetMeterProvider(mp interface{}) {
 panic("stub")
}

func defaultTracerValue() interface{} {
 panic("stub")
}

func defaultPropagatorsValue() interface{} {
 panic("stub")
}

func defaultMeterProvider() interface{} {
 panic("stub")
}

type tracerProvider struct {mtx interface{}; tracers interface{}; delegate interface{}}

type il struct {name interface{}; version interface{}}

type tracer struct {name interface{}; opts interface{}; provider interface{}; delegate interface{}}

type nonRecordingSpan struct {sc interface{}; tracer interface{}}

func (p *tracerProvider) setDelegate(provider interface{}) {
 panic("stub")
}

func (p *tracerProvider) Tracer(name interface{}, opts interface{}) interface{} {
 panic("stub")
}

func (t *tracer) setDelegate(provider interface{}) {
 panic("stub")
}

func (t *tracer) Start(ctx interface{}, name interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (s nonRecordingSpan) SpanContext() interface{} {
 panic("stub")
}

func (s nonRecordingSpan) TracerProvider() interface{} {
 panic("stub")
}

