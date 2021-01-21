package heartbeat

import (
	"fmt"
	"context"
	"github.com/AgentCoop/heartbeat/internal/api/heartbeat"
	"google.golang.org/grpc"
)

type server struct {
	heartbeat.UnimplementedHeartBeatServer
}

func (s *server) Foo(ctx context.Context, req *heartbeat.Request) (*heartbeat.Response, error) {
	fmt.Printf("foo request\n")
	return nil, nil
}

func RegisterServer(s *grpc.Server) {
	srv := &server{}
	heartbeat.RegisterHeartBeatServer(s, srv)
}
