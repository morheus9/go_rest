package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/morheus9/go_rest/internal/config"
	"github.com/morheus9/go_rest/internal/user"
	"github.com/morheus9/go_rest/internal/user/db"
	"github.com/morheus9/go_rest/pkg/client/mongodb"
	"github.com/morheus9/go_rest/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	cfgMongo := cfg.MongoDB
	mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	if err != nil {
		panic(err)
	}

	storage := db.NewStorage(mongoDBClient, cfgMongo.Collection, logger)

	user1 := user.User{
		ID:           "",
		Email:        "morheus12345@gmail.com",
		Username:     "morheus12345",
		PasswordHash: "password123",
	}

	user1ID, err := storage.Create(context.Background(), user1)
	if err != nil {
		panic(err)
	}
	logger.Info(user1ID)

	logger.Info("register new handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("create socket path")
		buildDir := path.Join("build")
		err := os.MkdirAll(buildDir, os.ModePerm) // Создаем папку build, если она не существует
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("create socket in ")
		socketPath := path.Join(buildDir, "app.sock")

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		if listenErr != nil {
			logger.Fatal(listenErr)
		}
		logger.Infof("server is listening on unix socket: %s", socketPath)

	} else {
		logger.Info("listen tcp port")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		if listenErr != nil {
			logger.Fatal(listenErr)
		}
		logger.Infof("server is listening on http://%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))

}
