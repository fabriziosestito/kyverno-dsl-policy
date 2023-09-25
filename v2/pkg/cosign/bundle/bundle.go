package bundle

import _ "unsafe"

type Embedme interface{}

type RekorBundle struct {SignedEntryTimestamp interface{}; Payload interface{}}

type RekorPayload struct {Body interface{}; IntegratedTime interface{}; LogIndex interface{}; LogID interface{}}

func EntryToBundle(entry interface{}) interface{} {
 panic("stub")
}

type RFC3161Timestamp struct {SignedRFC3161Timestamp interface{}}

func TimestampToRFC3161Timestamp(timestampRFC3161 interface{}) interface{} {
 panic("stub")
}

