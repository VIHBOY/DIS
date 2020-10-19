package main

import (
	"log"
	"net"
	"os"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

//CreateFile is
/***
* func CreateFile
**
* Crea archivos APPEND
**
* Input:
* string name : Nombre del archivo CSV
***/
func CreateFile(name string) {
	csvFile, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvFile.Close()
}

//Connect is
/***
* func connect
**
* Crear un server gRPC, para hacer conexiones a cliente y camion
**
* Input:
* net.Listener l : Credenciales para el servidor
***/
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
