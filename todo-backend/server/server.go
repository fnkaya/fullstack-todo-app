package server

import (
	"backend/handler"
	"backend/repository"
	"backend/service"
	"fmt"
	"net/http"
)

type server struct {
	Port int
}

func (s *server) StartServer() error {
	repo := repository.NewRepository()
	srv := service.NewService(repo)
	hndlr := handler.NewHandler(srv)
	http.HandleFunc("/api/todos", hndlr.HandleRequest)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
}

func NewServer(port int) *server {
	return &server{Port: port}
}
