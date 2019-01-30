package http10

import (
	"fmt"
	"net"
)

type Server struct {
	Addr string
	Port string
}

// Listener created server and listen
func (srv *Server) Listener() error {

	addr := srv.Addr + ":" + srv.Port
	ln, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println("err =", err)
		return err
	}
	defer ln.Close()

	fmt.Println("Listening on ", addr)
	fmt.Println("----------------------------------------------------------------------")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("err =", err)
			return err
		}

		go handleRequest(conn)

	}
}

func handleRequest(conn net.Conn) {

	// 1. Спарсить запрос и записать в Request
	// 2. Запустить Middlewares, если они определены
	// 3. Обработать запрос соответствующими методами (GET, POST, HEAD) если они определены
	// 4. Сгенерировать Responce
	// 5. Отправить ответ

	// 1. Парсим запрос и записываем в Request
	request := &Request{Header: make(Header)}
	request.init(conn)

	// Started Hadlers Here!
	M["/index"](request)

	fmt.Println("Method=" + request.Method)
	fmt.Println("URL=" + request.URL)
	fmt.Println("HTTP_v=" + request.HTTP_v)
	fmt.Println("User-Agent=" + request.Header.Get("User-Agent"))

	fmt.Println("**------------------------------------------------------------------------**")
	fmt.Println("End of responce")
	fmt.Println("**------------------------------------------------------------------------**")

	var res string = "HTTP/1.0 200 OK \n" +
		"Content-Type: text/plain; charset=utf-8\n" +
		"X-Content-Type-Options: application/json\n" +
		"Date: Thu, 24 Jan 2019 07:20:05 GMT\n" +
		"Connection: keep-alive\n" +
		"Content-Length: 32\n" +
		"\n" +
		"Hello, World! 1\n" +
		"Hello, World! 2\n"

	conn.Write([]byte(res))

	conn.Close()
}
