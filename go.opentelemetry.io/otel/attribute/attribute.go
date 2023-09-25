package attribute

import _ "unsafe"

type Embedme interface{}

type Encoder interface {}

type EncoderID struct {value interface{}}

type defaultAttrEncoder struct {pool interface{}}

func NewEncoderID() interface{} {
 panic("stub")
}

func DefaultEncoder() interface{} {
 panic("stub")
}

func (d *defaultAttrEncoder) Encode(iter interface{}) interface{} {
 panic("stub")
}

func copyAndEscape(buf interface{}, val interface{}) {
 panic("stub")
}

func (id EncoderID) Valid() interface{} {
 panic("stub")
}

type Filter func(interface{}) interface{}

func NewAllowKeysFilter(keys interface{}) interface{} {
 panic("stub")
}

func NewDenyKeysFilter(keys interface{}) interface{} {
 panic("stub")
}

type Iterator struct {storage interface{}; idx interface{}}

type MergeIterator struct {one interface{}; two interface{}; current interface{}}

type oneIterator struct {iter interface{}; done interface{}; attr interface{}}

func (i *Iterator) Next() interface{} {
 panic("stub")
}

func (i *Iterator) Label() interface{} {
 panic("stub")
}

func (i *Iterator) Attribute() interface{} {
 panic("stub")
}

func (i *Iterator) IndexedLabel() (interface{}, interface{}) {
 panic("stub")
}

func (i *Iterator) IndexedAttribute() (interface{}, interface{}) {
 panic("stub")
}

func (i *Iterator) Len() interface{} {
 panic("stub")
}

func (i *Iterator) ToSlice() interface{} {
 panic("stub")
}

func NewMergeIterator(s1, s2 interface{}) interface{} {
 panic("stub")
}

func makeOne(iter interface{}) interface{} {
 panic("stub")
}

func (oi *oneIterator) advance() {
 panic("stub")
}

func (m *MergeIterator) Next() interface{} {
 panic("stub")
}

func (m *MergeIterator) Label() interface{} {
 panic("stub")
}

func (m *MergeIterator) Attribute() interface{} {
 panic("stub")
}

type Key string

func (k Key) Bool(v interface{}) interface{} {
 panic("stub")
}

func (k Key) BoolSlice(v interface{}) interface{} {
 panic("stub")
}

func (k Key) Int(v interface{}) interface{} {
 panic("stub")
}

func (k Key) IntSlice(v interface{}) interface{} {
 panic("stub")
}

func (k Key) Int64(v interface{}) interface{} {
 panic("stub")
}

func (k Key) Int64Slice(v interface{}) interface{} {
 panic("stub")
}

func (k Key) Float64(v interface{}) interface{} {
 panic("stub")
}

func (k Key) Float64Slice(v interface{}) interface{} {
 panic("stub")
}

func (k Key) String(v interface{}) interface{} {
 panic("stub")
}

func (k Key) StringSlice(v interface{}) interface{} {
 panic("stub")
}

func (k Key) Defined() interface{} {
 panic("stub")
}

type KeyValue struct {Key interface{}; Value interface{}}

func (kv KeyValue) Valid() interface{} {
 panic("stub")
}

func Bool(k interface{}, v interface{}) interface{} {
 panic("stub")
}

func BoolSlice(k interface{}, v interface{}) interface{} {
 panic("stub")
}

func Int(k interface{}, v interface{}) interface{} {
 panic("stub")
}

func IntSlice(k interface{}, v interface{}) interface{} {
 panic("stub")
}

func Int64(k interface{}, v interface{}) interface{} {
 panic("stub")
}

func Int64Slice(k interface{}, v interface{}) interface{} {
 panic("stub")
}

func Float64(k interface{}, v interface{}) interface{} {
 panic("stub")
}

func Float64Slice(k interface{}, v interface{}) interface{} {
 panic("stub")
}

func String(k, v interface{}) interface{} {
 panic("stub")
}

func StringSlice(k interface{}, v interface{}) interface{} {
 panic("stub")
}

func Stringer(k interface{}, v interface{}) interface{} {
 panic("stub")
}

type Set struct {equivalent interface{}}

type Sortable []KeyValue

type Distinct struct {iface interface{}}

func EmptySet() interface{} {
 panic("stub")
}

func (d Distinct) reflectValue() interface{} {
 panic("stub")
}

func (d Distinct) Valid() interface{} {
 panic("stub")
}

func (l *Set) Len() interface{} {
 panic("stub")
}

func (l *Set) Get(idx interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (l *Set) Value(k interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (l *Set) HasValue(k interface{}) interface{} {
 panic("stub")
}

func (l *Set) Iter() interface{} {
 panic("stub")
}

func (l *Set) ToSlice() interface{} {
 panic("stub")
}

func (l *Set) Equivalent() interface{} {
 panic("stub")
}

func (l *Set) Equals(o interface{}) interface{} {
 panic("stub")
}

func (l *Set) Encoded(encoder interface{}) interface{} {
 panic("stub")
}

func empty() interface{} {
 panic("stub")
}

func NewSet(kvs interface{}) interface{} {
 panic("stub")
}

func NewSetWithSortable(kvs interface{}, tmp interface{}) interface{} {
 panic("stub")
}

func NewSetWithFiltered(kvs interface{}, filter interface{}) (interface{}, interface{}) {
 panic("stub")
}

func NewSetWithSortableFiltered(kvs interface{}, tmp interface{}, filter interface{}) (interface{}, interface{}) {
 panic("stub")
}

func filterSet(kvs interface{}, filter interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (l *Set) Filter(re interface{}) (interface{}, interface{}) {
 panic("stub")
}

func computeDistinct(kvs interface{}) interface{} {
 panic("stub")
}

func computeDistinctFixed(kvs interface{}) interface{} {
 panic("stub")
}

func computeDistinctReflect(kvs interface{}) interface{} {
 panic("stub")
}

func (l *Set) MarshalJSON() (interface{}, interface{}) {
 panic("stub")
}

func (l Set) MarshalLog() interface{} {
 panic("stub")
}

func (l *Sortable) Len() interface{} {
 panic("stub")
}

func (l *Sortable) Swap(i, j interface{}) {
 panic("stub")
}

func (l *Sortable) Less(i, j interface{}) interface{} {
 panic("stub")
}

func _() {
 panic("stub")
}

func (i Type) String() interface{} {
 panic("stub")
}

type unknownValueType struct {}

type Type int

type Value struct {vtype interface{}; numeric interface{}; stringly interface{}; slice interface{}}

func BoolValue(v interface{}) interface{} {
 panic("stub")
}

func BoolSliceValue(v interface{}) interface{} {
 panic("stub")
}

func IntValue(v interface{}) interface{} {
 panic("stub")
}

func IntSliceValue(v interface{}) interface{} {
 panic("stub")
}

func Int64Value(v interface{}) interface{} {
 panic("stub")
}

func Int64SliceValue(v interface{}) interface{} {
 panic("stub")
}

func Float64Value(v interface{}) interface{} {
 panic("stub")
}

func Float64SliceValue(v interface{}) interface{} {
 panic("stub")
}

func StringValue(v interface{}) interface{} {
 panic("stub")
}

func StringSliceValue(v interface{}) interface{} {
 panic("stub")
}

func (v Value) Type() interface{} {
 panic("stub")
}

func (v Value) AsBool() interface{} {
 panic("stub")
}

func (v Value) AsBoolSlice() interface{} {
 panic("stub")
}

func (v Value) asBoolSlice() interface{} {
 panic("stub")
}

func (v Value) AsInt64() interface{} {
 panic("stub")
}

func (v Value) AsInt64Slice() interface{} {
 panic("stub")
}

func (v Value) asInt64Slice() interface{} {
 panic("stub")
}

func (v Value) AsFloat64() interface{} {
 panic("stub")
}

func (v Value) AsFloat64Slice() interface{} {
 panic("stub")
}

func (v Value) asFloat64Slice() interface{} {
 panic("stub")
}

func (v Value) AsString() interface{} {
 panic("stub")
}

func (v Value) AsStringSlice() interface{} {
 panic("stub")
}

func (v Value) asStringSlice() interface{} {
 panic("stub")
}

func (v Value) AsInterface() interface{} {
 panic("stub")
}

func (v Value) Emit() interface{} {
 panic("stub")
}

func (v Value) MarshalJSON() (interface{}, interface{}) {
 panic("stub")
}

