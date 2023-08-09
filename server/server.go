package main

import (
	"google.golang.org/grpc"
	"hello_grpc/chat"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
		log.Fatalf("failed to listen: %s", err.Error())
	}

	s := grpc.NewServer()

	// the PB code we generated before providers a helper for registeration
	chat.RegisterChatServiceServer(s, &chat.Service{})

	// Under the hood it's just doing the following which seems like a better delegation of responsibility IMHO.
	// However, you are supposed to use the above canonically.
	// s.RegisterService(&chat.ChatService_ServiceDesc, s)

	// Start listening now that the service is registered.  This will block indefinitely. If you want to do something
	// else you'll need to wrap this in a goroutine.
	if err = s.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed to listen: %s", err.Error())
	}

}
