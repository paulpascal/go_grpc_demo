package main

import (
	"context"
	"fmt"
	"net"

	proto "grpc-eg-go/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Ping(ctx context.Context, request *proto.PingRequest) (response *proto.PongResponse, err error) {
	message, id := request.GetMessage(), request.GetId()
	return &proto.PongResponse{
		Message: fmt.Sprintf("Reply of %s with ID: %d is OK :)", message, id),
	}, nil
}
