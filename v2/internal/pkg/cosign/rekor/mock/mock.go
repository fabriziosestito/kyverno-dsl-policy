package mock

import _ "unsafe"

type Embedme interface{}

type EntriesClient struct {Entries interface{}}

func (m *EntriesClient) CreateLogEntry(_ interface{}, _ interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *EntriesClient) GetLogEntryByIndex(_ interface{}, _ interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *EntriesClient) GetLogEntryByUUID(params interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *EntriesClient) SearchLogQuery(params interface{}, opts interface{}) (interface{}, interface{}) {
 panic("stub")
}

func (m *EntriesClient) SetTransport(transport interface{}) {
 panic("stub")
}

