package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port int
	e    *echo.Echo
}

func NewServer(port int) *Server {
	server := &Server{
		port: port,
		e:    echo.New(),
	}

	// logger
	server.e.Use(middleware.Logger())

	return server
}

func (s *Server) Start() {
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf(":%d", s.port)))
}
