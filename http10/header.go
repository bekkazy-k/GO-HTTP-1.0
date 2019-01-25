package http10

import "net/textproto"

type Header map[string][]string

// Get метод возвращает заголовок по ключу
func (h Header) Get(key string) string {
	return textproto.MIMEHeader(h).Get(key)
}

// Set записывает заголовки.
// Заменяет значение, если заголовок с таким ключом существует
func (h Header) Set(key, value string) {
	textproto.MIMEHeader(h).Set(key, value)
}

// Add добавляет заголовки
func (h Header) Add(key, value string) {
	textproto.MIMEHeader(h).Add(key, value)
}
