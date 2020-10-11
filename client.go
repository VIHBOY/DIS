package main

import (
	"context"
	"log"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist26:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Holi soy el cliente",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("no se pudo DECIR HOLA: %s", err)
	}
	log.Printf("respuesta del server %s", response.Body)
}
