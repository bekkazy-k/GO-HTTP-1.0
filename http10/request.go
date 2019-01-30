package http10

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type Request struct {
	Method  string
	URL     string
	HTTP_v  string
	Header  Header
	Body    string
	GetBody func() (io.ReadCloser, error)
}

// init Парсит заголовки и записывает в Request
func (req *Request) init(conn net.Conn) {

	// 1. Спарсить начальную строку запроса
	// 2. Записать заголовки запроса //ReadMIMEHeader
	// 3. Записать тело запроса

	// TODO: Need Use Buffion Readline!
	// FIXME: Buffer size
	message := make([]byte, 1024)

	_, error := conn.Read(message)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	arr := strings.Split(string(message), "\n")

	for i, val := range arr {
		if val == "\r" {
			break
		}
		if i == 0 {
			// Начальная строка запроса
			stHeadLine := strings.Split(string(val), " ")
			req.Method = stHeadLine[0]
			req.URL = stHeadLine[1]
			req.HTTP_v = stHeadLine[2]
		} else {
			// Заголовки
			arr := strings.Split(val, ": ")
			req.Header.Add(arr[0], arr[1])
		}
	}

}
