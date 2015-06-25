package core

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	Server struct {
		Address string
		Router  *httprouter.Router
	}
)

func NewServer(address string) *Server {
	return &Server{
		Address: address,
		Router:  httprouter.New(),
	}
}

func (s Server) DELETE(path string, handle httprouter.Handle) {
	s.Router.DELETE(path, handle)
}

func (s Server) GET(path string, handle httprouter.Handle) {
	s.Router.GET(path, handle)
}

func (s Server) POST(path string, handle httprouter.Handle) {
	s.Router.POST(path, handle)
}

func (s Server) Run() {
	log.Fatal(http.ListenAndServe("localhost:8080", s.Router))
}
