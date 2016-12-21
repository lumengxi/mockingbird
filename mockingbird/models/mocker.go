package models

import (
	"strings"
)

type Mocker struct {
	ID string `json:"uuid"`
	Status int `json:"status"`
	ContentType string `json:"contentType"`
	Charset string `json:"charset"`
	Location string `json:"location"`
	Body string `json:"body"`
	Headers map[string]string `json:"headers"`
}

func (m *Mocker) MakeHeaders() (map[string]string) {
	m.Headers["Content-Type"] = m.ContentType + "; charset=" + strings.ToLower(m.Charset)
	m.Headers["Location"] = m.Location

	return m.Headers
}
