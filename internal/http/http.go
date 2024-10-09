package http

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port   int
	Router *Router
}

type Handlers interface {
	RegisterRoutes(router *Router)
}

func Init(port int) *Server {
	return &Server{
		Port: port,
		Router: &Router{
			Routes: []Route{},
		},
	}
}

func (server *Server) Register(handlers Handlers) {
	handlers.RegisterRoutes(server.Router)
}

func (server *Server) Run() {
	log.Printf("🚀 Server ready at http://localhost:%d\n", server.Port)

	var err error = Listen(server.Port, server.Router)

	if err != nil {
		log.Fatalf("Failed to run the server, %v", err)
	}
}

func Listen(port int, router *Router) error {
	var address string = fmt.Sprintf(":%v", port)
	return http.ListenAndServe(address, router)
}
