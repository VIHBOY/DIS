package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

type Camion struct {
	id          int
	Nombre      string
	Paquete1    *chat.Paquete
	Paquete2    *chat.Paquete
	LleveRetail string
}

func EnviarPaquete(camion Camion, c chat.ChatServiceClient, np int) {
	rand.Seed(time.Now().UnixNano())
	prob := rand.Intn(6-1) + 1

	if np == 1 {
		if prob == 5 {
			me3 := chat.Message{
				Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
			}
			camion.Paquete1.Intentos++
			camion.Paquete1.Estado = "No Recibido"
			c.CambiarEstado(context.Background(), &me3)
			c.CambiarIntentos(context.Background(), &me3)
		} else {
			me3 := chat.Message{
				Body: camion.Paquete1.GetTrack() + "%" + "Recibido",
			}
			camion.Paquete1.Intentos++
			camion.Paquete1.Estado = "Recibido"
			c.CambiarEstado(context.Background(), &me3)
			c.CambiarIntentos(context.Background(), &me3)
		}
	}
	if np == 2 {
		if prob == 5 {
			me3 := chat.Message{
				Body: camion.Paquete2.GetTrack() + "%" + "No Recibido",
			}
			camion.Paquete2.Intentos++
			camion.Paquete2.Estado = "No Recibido"
			c.CambiarEstado(context.Background(), &me3)
		} else {
			me3 := chat.Message{
				Body: camion.Paquete2.GetTrack() + "%" + "Recibido",
			}
			camion.Paquete2.Intentos++
			camion.Paquete2.Estado = "Recibido"
			c.CambiarEstado(context.Background(), &me3)
		}
	}

}

//NewOrden is
func Send(camion Camion) {
	var conn *grpc.ClientConn
	conn, _ = grpc.Dial(":9000", grpc.WithInsecure())
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: camion.Nombre,
	}
	var can sync.Mutex
	can.Lock()
	p1, _ := c.Recibir2(context.Background(), &message)
	can.Unlock()
	can.Lock()
	p2, _ := c.Recibir2(context.Background(), &message)
	can.Unlock()

	if p1.GetId() != "NOHAY" {
		if p2.GetId() != "NOHAY" {
			camion.Paquete1 = p1
			camion.Paquete2 = p2
			fmt.Println("Camion %d recibio %d,", camion.id, camion.Paquete1.GetId())
			fmt.Println("Camion %d recibio %d,", camion.id, camion.Paquete2.GetId())
			me := chat.Message{
				Body: camion.Paquete1.GetTrack() + "%" + "En Camino",
			}
			c.CambiarEstado(context.Background(), &me)
			me = chat.Message{
				Body: camion.Paquete2.GetTrack() + "%" + "En Camino",
			}
			c.CambiarEstado(context.Background(), &me)
			if camion.Paquete1.GetValor() >= camion.Paquete2.GetValor() {
				can.Lock()
				EnviarPaquete(camion, c, 1)
				can.Unlock()
				/////////////////////////////////
				can.Lock()
				EnviarPaquete(camion, c, 2)
				can.Unlock()
			} else {
				can.Lock()
				EnviarPaquete(camion, c, 2)
				can.Unlock()
				/////////////////////////////////
				can.Lock()
				EnviarPaquete(camion, c, 1)
				can.Unlock()
			}
		} else {
			camion.Paquete1 = p1
			fmt.Println("Camion %d recibio %d,", camion.id, camion.Paquete1.GetId())
			me := chat.Message{
				Body: camion.Paquete1.GetTrack() + "%" + "En Camino",
			}
			c.CambiarEstado(context.Background(), &me)
			can.Lock()
			EnviarPaquete(camion, c, 1)
			can.Unlock()
		}
	}
}
func main() {

	CamionRetail1 := Camion{
		id:     1,
		Nombre: "Retail",
	}
	CamionRetail2 := Camion{
		id:     2,
		Nombre: "Retail",
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese Tiempo espera Camion")
	tespera, _ := reader.ReadString('\n')
	fmt.Println("Ingrese Tiempo de envio Paquete")
	tenvio, _ := reader.ReadString('\n')
	log.Printf("Tiempo espera Camion: %s", tespera)
	log.Printf("Tiempo de envio Camion: %s", tenvio)
	var can sync.Mutex

	var2 := time.Duration(5) * time.Second
	for {
		time.Sleep(var2)
		can.Lock()
		go Send(CamionRetail1)
		can.Unlock()
		can.Lock()
		go Send(CamionRetail2)
		can.Unlock()
	}

}
