package http10

import "net/textproto"

// Header описывает структуру заголовков
type Header map[string][]string

// Get метод возвращает заголовок по ключу
func (h Header) Get(key string) string {
	return textproto.MIMEHeader(h).Get(key)
}
