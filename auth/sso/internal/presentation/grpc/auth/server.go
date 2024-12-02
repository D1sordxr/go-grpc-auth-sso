package auth

import (
	"fmt"
	services "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

type Server struct {
	Service services.AuthServer
	Server  *grpc.Server
}

func NewGRPCServer() *Server {
	return &Server{
		Server: grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		})),
	}
}

func (s *Server) Run(port int) error {
	grpcServer := s.Server
	services.RegisterAuthServer(grpcServer, s.Service)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to start TCP listener on port %v: %v", port, err)
	}

	reflection.Register(grpcServer)
	s.Server = grpcServer
	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Down() {
	s.Server.GracefulStop()
}
