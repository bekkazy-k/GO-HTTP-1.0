package http10

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Server описывает структуру сервера
type Server struct {
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
	b := bufio.NewReader(conn)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // for example EOF
			break
		}
		fmt.Print(string(line))
		conn.Write(line)
	}
}
