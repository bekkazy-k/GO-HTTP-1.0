package http10

import (
	"fmt"
	"net"
	"os"
)

// Server описывает структуру сервера
type Server struct {
	Addr   string
	Port   string
	Status string
}

// CreateServer функция создает новый экземпляр сервера
func CreateServer() *Server {
	return &Server{}
}

// ChangePort метод меняет порт экземпляра сервера
func (srv *Server) ChangePort(port string) {
	srv.Port = port
}

// Listen метод создает новое соединение
// Принимает хост, порт и тип соединения
func (srv *Server) Listen(connHost, connPort, connType string) {
	// Слушает входящие сообщения
	l, err := net.Listen(connType, connHost+":"+connPort)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Закрывает Listener когда приложение закрывается
	defer l.Close()
	fmt.Println("Listening on " + connHost + ":" + connPort)
	fmt.Println("----------------------------------------------------------------------")

	for {
		// Прослушивает входящие сообщения
		conn, err := l.Accept()

		// addr := l.Addr()
		// netw := addr.Network()
		// str := addr.String()

		// fmt.Println("netw:", netw)
		// fmt.Println("str:", str)

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Обрабатывает содинения в новом goroutine
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	message := make([]byte, 1024)

	_, error := conn.Read(message)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	fmt.Println(string(message))

	req := &Request{Header: make(Header)}
	// req.Header = make(map[string]string)

	req.ParseInnerData(string(message))

	fmt.Println("Headers host =", req.Header.Get("Host"))
	// req.make(header)

	fmt.Println("**------------------------------------------------------------------------**")
	fmt.Println("End of responce")
	fmt.Println("**------------------------------------------------------------------------**")

	conn.Write([]byte("Responce typed here"))

	conn.Close()
}
