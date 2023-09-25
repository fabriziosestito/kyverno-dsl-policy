package codes

import _ "unsafe"

type Embedme interface{}

type Code uint32

func (c Code) String() interface{} {
 panic("stub")
}

func (c *Code) UnmarshalJSON(b interface{}) interface{} {
 panic("stub")
}

func (c *Code) MarshalJSON() (interface{}, interface{}) {
 panic("stub")
}

