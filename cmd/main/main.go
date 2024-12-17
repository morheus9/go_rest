package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/morheus9/go_rest/internal/user"
)

func main() {
	fmt.Println("create router")
	router := httprouter.New()

	fmt.Println("register new handler")
	handler := user.NewHandler()
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Listening on http://%s\n", listener.Addr())
	log.Fatalln(server.Serve(listener))

}
