package webserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerImpl struct {
	port   int
	router *mux.Router
}

type Server interface {
	Initialize() error
	Start() error
}

func NewServer(port int) Server {
	return &ServerImpl{port: port}
}

func (s *ServerImpl) Start() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%v", s.port), s.router)
}

func (s *ServerImpl) Initialize() error {
	s.router = mux.NewRouter()
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }).Methods("GET")
	s.router.HandleFunc("/compile", compileHandler).Methods("POST")
	return nil
}
