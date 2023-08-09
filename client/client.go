package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hello_grpc/chat"
	"log"
)

func main() {

	// make the connection
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %s", err.Error())
	}

	// service client, this gets generated for you by protoc go-grpc-out
	c := chat.NewChatServiceClient(conn)
	r, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello from the client!"})
	if err != nil {
		log.Fatalf("Could not say hello: %s", err.Error())
	}
	log.Println(r.Body)

}
