package main

import (
	"log"
	"net"
	"os"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

//CreateFile is
func CreateFile(name string) {
	csvFile, err := os.Create(name)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvFile.Close()
}

func main() {
	CreateFile("dblogistica.csv")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed %v", err)
	}
	s := chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
