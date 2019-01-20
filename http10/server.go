package http10

import (
	"fmt"
	"net"
	"os"
	"strings"
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
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error(), reqLen)
	}
	fmt.Println("***** ReqBody Start ***********************")
	reqBody := string(buf)
	fmt.Println(reqBody)
	fmt.Println("***** ReqBody End ***********************")

	fmt.Println("***** Split ReqBody Start ***********************")
	a := strings.Split(reqBody, "\n")
	fmt.Printf("%q\n", a)
	fmt.Println("-----------------------------------------------------")
	fmt.Println(a)
	fmt.Println("***** Split ReqBody End ***********************")

	for i, v := range a {
		if v == "\n" {
			fmt.Println(i, " = ", v, "is new line/\n")
		} else if v == " " {
			fmt.Println(i, " = ", v, "is space")
		} else if v == "\r" {
			fmt.Println(i, " = ", v, "is RRRRR")
		} else {
			fmt.Println(i, " = ", v)
		}
	}

	fmt.Println("End of responce")

	conn.Write([]byte("Responce typed here"))
	conn.Close()
}
