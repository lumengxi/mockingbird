package main

import (
	"strings"
)

type MockerConfig struct {
	ContentType  string `json:"contentType"`
	Charset      string `json:"charset"`
	Location     string `json:"location"`
	Body         string `json:"body"`
	ExtraHeaders map[string]string `json:"headers"`
}


type Mocker struct {
	ID string `json:"uuid"`
	Status bool `json:"status"`
	Name string `json:"name"`
	MockerConfig MockerConfig `json:"config"`
}

func (m *MockerConfig) MakeHeaders() (map[string]string) {
	m.ExtraHeaders["Content-Type"] = m.ContentType + "; charset=" + strings.ToLower(m.Charset)
	m.ExtraHeaders["Location"] = m.Location

	return m.ExtraHeaders
}
