package main

import (
    "fmt"
    "log"
    "net"
)

const PORT = 3333

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", PORT) )
    if err != nil {
        errMessage := fmt.Sprintf("Error binding to port %d: %s", PORT, err)
        log.Fatal(errMessage)
    }
	defer l.Close()

	log.Printf("Server listening on port %d", PORT)
	
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading the request: %v", err)
		return
	}

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	log.Printf("Handled connection from %s", conn.RemoteAddr())
}
