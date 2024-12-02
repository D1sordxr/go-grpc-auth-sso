package auth

import (
	services "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"time"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

type Server struct {
	//Service services.AuthServer
	Service services.UnimplementedAuthServer
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

func (s *Server) RegisterServer() error {
	grpcServer := s.Server
	services.RegisterAuthServer(grpcServer, s.Service)

	reflection.Register(grpcServer)
	return nil
}

func (s *Server) Down() {
	s.Server.GracefulStop()
}
