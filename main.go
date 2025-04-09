package main

import (
	"errors"
	"fmt"
	"net/http"
	"io"
	"context"
	"net"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	
	ctx := context.Background()
	server := &http.Server{
		Addr: "localhost:3000",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, "keyServerAddr", l.Addr().String())
			return ctx
		},
	}
	
	fmt.Println("listening on port 3000")

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	hasSecret := r.URL.Query().Has("secret")

	fmt.Printf(
		"get / request. first(%t)=%s, second(%t)=%s\n",
		hasFirst, first,
		hasSecond, second,
	)

	io.WriteString(w, "This is my server!\n")
	if(hasSecret) {
		io.WriteString(w, "You have found the secret")
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get /hello request")
	io.WriteString(w, "Hello from Brasil!\n")
}

