package http10

// Listen Создает новое соединение
func Listen(port string) error {
	if port == "" {
		port = "80"
	}
	server := &Server{Addr: "localhost", Port: port}
	return server.Listener()
}
