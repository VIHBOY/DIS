package chat

import (
	context "context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Rece mensaje form client: %s", message.Body)
	return &Message{Body: "Golla de server"}, nil
}
