package main

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type GrpcServer struct {
	config      *Config
	grpcApi     *GrpcApi
	_grpcServer *grpc.Server
}

func NewGrpcServer(config *Config, grpcApi *GrpcApi) *GrpcServer {
	return &GrpcServer{
		config:  config,
		grpcApi: grpcApi,
	}
}

func (s *GrpcServer) Run() {
	lis, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		s.config.Log.Log.Fatal(err)
	}
	opt := []grpc.ServerOption{}
	s._grpcServer = grpc.NewServer(opt...)
	RegisterIncrementerServer(s._grpcServer, s.grpcApi)
	err = s._grpcServer.Serve(lis)
	if err != nil {
		s.config.Log.Log.Fatal(err)
	}
}

func (s *GrpcServer) ReadyToStop() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM) // Отлавливаем в канал interrupt сигналы os.Interrupt и syscall.SIGTERM

	<-interrupt // Здесь исполнение кода блокируется, пока не не будет получен сигнал ОС

	s.config.Log.Log.Info("Stopping server...")

	s._grpcServer.GracefulStop()
}
