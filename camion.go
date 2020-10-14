package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

type Cola struct {
	Nombre string
	Cola   []string
}

type Camion struct {
	Nombre      string
	Paquete1    string
	Paquete2    string
	LlevoRetail string
}

//NewOrden is

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	if err != nil {
		log.Fatalf("no se pudo DECIR HOLA: %s", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese Tiempo espera Camion")
	tespera, _ := reader.ReadString('\n')
	fmt.Println("Ingrese Tiempo de envio Paquete")
	tenvio, _ := reader.ReadString('\n')
	log.Printf("Tiempo espera Camion: %s", tespera)
	log.Printf("Tiempo de envio Camion: %s", tenvio)
	message := chat.Message{
		Body: "%",
	}
	var2 := time.Duration(5) * time.Second
	for {
		time.Sleep(var2)
		response, _ := c.Recibir(context.Background(), &message)
		log.Printf("Su codigo de tracking es %s", response.Body)
		time.Sleep(var2)
		response, _ = c.Recibir(context.Background(), &message)
		log.Printf("Su codigo de tracking es %s", response.Body)
	}

}
