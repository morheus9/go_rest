package main

import (
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/morheus9/go_rest/internal/user"
	"github.com/morheus9/go_rest/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register new handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()

	logger.Info("start application")

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Listening on http://", listener.Addr())
	logger.Fatal(server.Serve(listener))

}
