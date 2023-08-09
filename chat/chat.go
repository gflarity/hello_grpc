package chat

import (
	"context"
	"log"
)

type Service struct {

	// This is a requirement of the interface required for registration below.
	UnimplementedChatServiceServer
}

func (s *Service) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Message Resceived From Client: %s", message.Body)
	return &Message{Body: "Hello from the server!"}, nil
}
