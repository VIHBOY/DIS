package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

type Camion struct {
	Nombre      string
	Paquete1    *chat.Paquete
	Paquete2    *chat.Paquete
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
	if p1.GetId() != "NOHAY" {
		if p2.GetId() != "NOHAY" {
			camion.Paquete1 = p1
			camion.Paquete2 = p2
			fmt.Println(camion.Paquete1.GetId())
			fmt.Println(camion.Paquete2.GetId())
			me := chat.Message{
				Body: camion.Paquete1.GetTrack() + "%" + "En Camino",
			}
			c.CambiarEstado(context.Background(), &me)
			if camion.Paquete1.GetValor() >= camion.Paquete2.GetValor() {
				me := chat.Message{
					Body: camion.Paquete1.GetTrack() + "%" + "En Camino",
				}
				c.CambiarIntentos(context.Background(), &me)
				rand.Seed(time.Now().UnixNano())
				prob := rand.Intn(6-1) + 1
				if prob == 5 {
					me := chat.Message{
						Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
					}
					c.CambiarEstado(context.Background(), &me)
				} else {
					me := chat.Message{
						Body: camion.Paquete1.GetTrack() + "%" + "Recibido",
					}
					c.CambiarEstado(context.Background(), &me)
				}
			}
		}
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
