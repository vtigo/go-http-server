package main

import (
    "fmt"
    "log"
    "net"
)

const PORT = 3333

func main() {
    conn := startServer(PORT)
    defer conn.Close()

    buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading the request: %v", err)
		return
	}
    
    conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}

func startServer(port int) net.Conn {
    l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
    if err != nil {
        errMessage := fmt.Sprintf("Error binding to port %d: %s", port, err)
        log.Fatal(errMessage)
    }
    conn, err := l.Accept()
    if err != nil {
        errMessage := fmt.Sprintf("Error accepting connection: %s", err)
        log.Fatal(errMessage)
    }
    return conn
}
