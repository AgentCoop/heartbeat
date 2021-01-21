package api

import (
	"fmt"
	"context"
	"heartbeat.service/internal/api/heartbeat"
	"google.golang.org/grpc"
)

func Foo() {
	fmt.Println("Hello, World!")
}

type server struct {
	heartbeat.UnimplementedHeartBeatServer
}

func (s *server) Foo(ctx context.Context, req *heartbeat.Request) (*heartbeat.Response, error) {
	fmt.Printf("foo request\n")
	return nil, nil
}

func NewRequest() *heartbeat.Request {
	req := &heartbeat.Request{}
	return req
}

type Client interface {
	heartbeat.HeartBeatClient
}
//
func NewClient(conn grpc.ClientConnInterface) Client {
	return heartbeat.NewHeartBeatClient(conn)
}

func RegisterServer(s *grpc.Server) {
	srv := &server{}
	heartbeat.RegisterHeartBeatServer(s, srv)
}
