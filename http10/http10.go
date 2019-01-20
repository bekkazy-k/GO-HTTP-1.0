package http10

import (
	"fmt"
)

func Start() {
	fmt.Println("Package started!")
	srv := CreateServer()
	srv.Listen("localhost", "5000", "tcp")
}
