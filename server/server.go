package server

import (
	"github.com/DimitryEf/incrementer-api/api"
	"github.com/DimitryEf/incrementer-api/config"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
)

// Server - структура gRPC-сервера инкрементора
type Server struct {
	config *config.Config
	api    *api.Api
	server *grpc.Server
}

// NewServer - конструктор gRPC-сервера
func NewServer(config *config.Config, api *api.Api) *Server {
	return &Server{
		config: config,
		api:    api,
	}
}

// Run - запуск сервера
func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		s.config.Logger.Log.Fatal(err)
	}

	// Интерцептор для middleware логирования сервера
	s.server = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(
					logrus.NewEntry(s.config.Logger.Log),
				))))

	// Регистрация АПИ сервера
	api.RegisterIncrementerServer(s.server, s.api)

	// Запуск сервера
	s.config.Logger.Log.Infof("Starting server on %s...", s.config.Port)
	err = s.server.Serve(lis)
	if err != nil {
		s.config.Logger.Log.Fatal(err)
	}
}

// ReadyToStop - отслеживание сигналов OS для Graceful shutdown
func (s *Server) ReadyToStop() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM) // Отлавливаем в канал interrupt сигналы os.Interrupt и syscall.SIGTERM

	<-interrupt // Здесь исполнение кода блокируется, пока не не будет получен сигнал ОС

	s.config.Logger.Log.Info("Stopping server...")

	s.server.GracefulStop()
}
