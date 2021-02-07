package main

import (
	"context"
	"fmt"
	"net"

	"github.com/rahmancam/grpc-tuts/echo"
	"google.golang.org/grpc"
)

// EchoServer server
type EchoServer struct{}

// Echo service for Echo server
func (e *EchoServer) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Response: "Echo: " + req.Message,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	srv := &EchoServer{}
	echo.RegisterEchoServerServer(s, srv)
	fmt.Println("Echo server serving @ port 8080")
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
