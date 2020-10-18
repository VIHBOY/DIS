package main

import (
	"log"
	"net"
	"os"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

//Server is
type Server struct {
}

//CreateFile is
func CreateFile(name string) {
	csvFile, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvFile.Close()
}

func connect(l net.Listener) {
	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

func main() {

	CreateFile("dblogistica.csv")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	connect(lis)
}
