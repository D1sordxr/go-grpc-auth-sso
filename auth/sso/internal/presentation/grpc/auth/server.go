package auth

//import (
//	"context"
//	ssov1 "github.com/D1sordxr/aviasales/auth/protos/gen/go/d1sx.sso.v1"
//	ssoV1 "github.com/D1sordxr/aviasales/auth/protos/gen/go/sso"
//	"google.golang.org/grpc"
//)
//
//type ServerAPI struct {
//	ssov1.UnimplementedAuthServer
//}
//
//func Register(gRPC *grpc.Server) {
//	ssov1.RegisterAuthServer(gRPC, &ServerAPI{})
//}
//
//func (s *ServerAPI) Login(ctx context.Context, req ssoV1.LoginRequest) (*ssoV1.LoginResponse, error) {
//	panic("implement me")
//}
