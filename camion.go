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

type Camion struct {
	Nombre      string
	Paquete1    string
	Paquete2    string
	LleveRetail string
}

//NewOrden is

func main() {

	CamionRetail1 := Camion{
		Nombre: "Retail",
	}
	CamionRetail2 := Camion{
		Nombre: "Retail",
	}
	CamionNormal := Camion{
		Nombre: "Normal",
	}
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

	var2 := time.Duration(5) * time.Second
	for {
		message := chat.Message{
			Body: CamionRetail1.Nombre,
		}
		time.Sleep(var2)
		response, _ := c.Recibir(context.Background(), &message)
		fmt.Println(response.Body)
		message = chat.Message{
			Body: CamionRetail2.Nombre,
		}
		time.Sleep(var2)
		response, _ = c.Recibir(context.Background(), &message)
		fmt.Println(response.Body)
		message = chat.Message{
			Body: CamionNormal.Nombre,
		}
		time.Sleep(var2)
		response, _ = c.Recibir(context.Background(), &message)
		fmt.Println(response.Body)

	}

}
