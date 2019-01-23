package http10

import (
	"strings"
)

// Request описывает структуру запроса
type Request struct {
	// Method (GET, POST, PUT, DELETE и т.д)
	Method string

	// Header map[string]string
	Header Header

	// Body   Body
}

// ParseInnerData разделяет Header запроса от Body
// И записывает их в соответствующие экземпляры структур
func (r *Request) ParseInnerData(str string) {
	a := strings.Split(str, "\n")
	for i, val := range a {
		if val == "\r" {
			break
		}
		if i > 0 {
			a := strings.Split(val, ": ")
			r.Header.Add(a[0], a[1])
		}
	}
}
