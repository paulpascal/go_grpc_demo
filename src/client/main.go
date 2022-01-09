package main

import (
	"context"
	"fmt"

	"grpc-eg-go/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := proto.NewServiceClient(conn)
	request := &proto.PingRequest{
		Message: "Ping me",
		Id:      uint64(20),
	}
	ctx := context.Background()
	response, err := client.Ping(ctx, request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Message)
}
