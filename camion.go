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
	Paquete1    chat.Paquete
	Paquete2    chat.Paquete
	LleveRetail string
}

//NewOrden is
func Send(camion Camion) {
	var conn *grpc.ClientConn
	conn, _ = grpc.Dial(":9000", grpc.WithInsecure())
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: camion.Nombre,
	}
	p1, _ := c.Recibir2(context.Background(), &message)
	p2, _ := c.Recibir2(context.Background(), &message)
	camion.Paquete1 = chat.Paquete{
		Id:       p1.GetId(),
		Track:    p1.GetTrack(),
		Tipo:     p1.GetTipo(),
		Intentos: p1.GetIntentos(),
		Estado:   p1.GetEstado(),
	}
	camion.Paquete2 = chat.Paquete{
		Id:       p2.GetId(),
		Track:    p2.GetTrack(),
		Tipo:     p2.GetTipo(),
		Intentos: p2.GetIntentos(),
		Estado:   p2.GetEstado(),
	}

}
func main() {

	CamionRetail1 := Camion{
		Nombre: "Retail",
	}
	CamionRetail2 := Camion{
		Nombre: "Retail",
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
		time.Sleep(var2)
		go Send(CamionRetail1)
		go Send(CamionRetail2)

	}

}
