package server

import (
	"fmt"
	"net/http"
	"os"
	"project2-microservice-go/database"
	"project2-microservice-go/internal/notification-service/routers"
	"strconv"
	"time"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("NOTIFICATION_SERVICE_PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	router := routers.NewRouter()

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      router.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
