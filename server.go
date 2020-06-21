package main

import (
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

type Server struct {
	Config     *Config
	Api        *Api
	Mid        *Middleware
	HttpServer *http.Server
}

func NewServer(config *Config, api *Api, mid *Middleware) *Server {
	return &Server{
		Config: config,
		Api:    api,
		Mid:    mid,
		HttpServer: &http.Server{
			Addr:           net.JoinHostPort(config.Host, config.Port),
			ReadTimeout:    config.ReadTimeout,
			WriteTimeout:   config.WriteTimeout,
			MaxHeaderBytes: config.MaxHeaderBytes,
		},
	}
}

func (server *Server) Handler() http.Handler {

	r := mux.NewRouter()

	// Устанавливаем единственную хэндлер-функцию
	r.HandleFunc("/api", server.Api.Do).Methods(http.MethodPost)

	r.Use(server.Mid.AddLogger())
	r.Use(server.Mid.AddHeaders())

	return r
}

func (server *Server) Run() {

	server.HttpServer.Handler = server.Handler()

	err := server.HttpServer.ListenAndServe()
	if err != nil {
		server.Config.Log.Log.Fatal(err)
	}
}
