package main

import (
	"net/http"
)

type MockerConfig struct {
	StatusCode   int `json:"status_code"`
	ContentType  string `json:"content_type"`
	Charset      string `json:"charset"`
	Body         string `json:"body"`
	ExtraHeaders map[string]string `json:"headers"`
}


type Mocker struct {
	ID string `json:"uuid"`
	Status bool `json:"status"`
	Name string `json:"name"`
	MockerConfig MockerConfig `json:"config"`
}

func (m *MockerConfig) MakeHeaders() http.Header {

	if m.ExtraHeaders == nil {
		m.ExtraHeaders = make(map[string]string)
	}

	if len(m.ContentType) == 0 {
		m.ContentType = "text/plain"
	}

	if len(m.Charset) == 0 {
		m.Charset = "us-ascii"
	}

	m.ExtraHeaders["Content-Type"] = m.ContentType + "; charset=" + m.Charset

	header := http.Header{}
	for headerName, headerValue := range m.ExtraHeaders {
		header.Add(headerName, headerValue)
	}

	return header
}
