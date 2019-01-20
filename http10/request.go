package http10

// Request описывает структуру запроса
type Request struct {
	// Method (GET, POST, PUT, DELETE и т.д)
	Method string

	Header Header

	// Body   Body
}
