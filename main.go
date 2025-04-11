package main

import (
	"fmt"
	"log"
	"net"
)

const PORT = 3333

func main() {
	conn := startServer(PORT)

	buffer := make([]byte, 1024)
	conn.Read(buffer)
	request := string(buffer)

	fmt.Println(request)

	conn.Write([]byte("HTTP/1.1 200 OK"))
}

func startServer(port int) net.Conn {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		errMessage := fmt.Sprintf("Error binding to port %d: %s", port, err )
		log.Fatal(errMessage)
	}

	conn, err := l.Accept()
	if err != nil {
		errMessage := fmt.Sprintf("Error accepting connection: %d", err )
		log.Fatal(errMessage)
	}

	return conn
}


