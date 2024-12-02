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
	service services.UnimplementedAuthServer
	server  *grpc.Server
}

func NewGrpcServer(service services.UnimplementedAuthServer) *Server {
	return &Server{
		service: service,
	}
}

func (g *Server) Run(port string) error {
	s := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: maxConnectionIdle * time.Minute,
		Timeout:           gRPCTimeout * time.Second,
		MaxConnectionAge:  maxConnectionAge * time.Minute,
		Time:              gRPCTime * time.Minute,
	}))
	services.RegisterAuthServer(s, g.service)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	reflection.Register(s)
	g.server = s
	err = s.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
func (g *Server) Down() {
	g.server.GracefulStop()
}
