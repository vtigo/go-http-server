package main

import (
	"fmt"
	"log"
	"net"
	"strings"
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
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading the request: %v", err)
		return
	}
	requestString := string(buffer[:n])
	requestLine, requestHeaders, requestBody := parseRequest(requestString)
	
	fmt.Printf("line: %s\n", requestLine)
	fmt.Printf("headers: %v\n", requestHeaders)
	fmt.Printf("body: %s\n", requestBody)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	log.Printf("Handled connection from %s", conn.RemoteAddr())
}

func parseRequest(request string) (string, map[string]string, string) {
	parts := strings.SplitN(request, "\r\n\r\n", 2)
	headerLines := strings.Split(parts[0], "\r\n")
	body := ""
	if len(parts[1]) > 1 {
		body = parts[1]
	}

	headers := make(map[string]string)
	for i:=1; i < len(headerLines); i++ {
		if headerLines[i] == " " {
			continue
		}
		parts := strings.SplitN(headerLines[i], ": ", 2)
		headers[parts[0]] = parts[1]
	}

	return string(headerLines[0]), headers, body
}
