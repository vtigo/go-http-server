package main

import (
	"fmt"
	"errors"
	"net/http"
	"os"
	"io"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe("127.0.0.1:3000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	io.WriteString(w, "This is my server!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /hello request")
}

