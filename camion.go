package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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

func EnviarPaquete(camion *Camion, c chat.ChatServiceClient, np int) {
	rand.Seed(time.Now().UnixNano())
	prob := rand.Intn(6-1) + 1

	if np == 1 {
		if prob == 5 {
			me3 := chat.Message{
				Body: camion.Paquete1.GetTrack() + "%" + "En Camino",
			}
			camion.Paquete1.Estado = "En Camino"
			camion.Paquete1.Intentos++
			log.Printf("Camion: %d: Paquete: %s Intentos %d:\n", camion.id, camion.Paquete1.Id, camion.Paquete1.Intentos)
			c.CambiarEstado(context.Background(), &me3)

		} else {
			me3 := chat.Message{
				Body: camion.Paquete1.GetTrack() + "%" + "Recibido",
			}
			camion.Paquete1.Estado = "Recibido"
			camion.Paquete1.Intentos++
			log.Printf("Camion: %d: Paquete: %s Intentos %d:\n", camion.id, camion.Paquete1.Id, camion.Paquete1.Intentos)
			c.CambiarEstado(context.Background(), &me3)
		}
	}
	if np == 2 {
		if prob == 5 {
			me3 := chat.Message{
				Body: camion.Paquete2.GetTrack() + "%" + "En Camino",
			}
			camion.Paquete2.Estado = "En Camino"
			camion.Paquete2.Intentos++
			log.Printf("Camion: %d: Paquete: %s Intentos %d:\n", camion.id, camion.Paquete2.Id, camion.Paquete2.Intentos)
			c.CambiarEstado(context.Background(), &me3)
		} else {
			me3 := chat.Message{
				Body: camion.Paquete2.GetTrack() + "%" + "Recibido",
			}
			camion.Paquete2.Estado = "Recibido"
			camion.Paquete2.Intentos++
			log.Printf("Camion: %d: Paquete: %s Intentos %d:\n", camion.id, camion.Paquete2.Id, camion.Paquete2.Intentos)
			c.CambiarEstado(context.Background(), &me3)
		}
	}

}

//NewOrden is
func Send(camion Camion, tespera int, tenvio int) {
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
			log.Printf("Camion %d recibio %s,", camion.id, camion.Paquete1.GetId())
			log.Printf("Camion %d recibio %s,", camion.id, camion.Paquete2.GetId())
			if camion.Paquete1.GetValor() >= camion.Paquete2.GetValor() {
				EnviarPaquete(&camion, c, 1)
				/////////////////////////////////
				EnviarPaquete(&camion, c, 2)

				if camion.Paquete1.Intentos < 3 && camion.Paquete1.GetEstado() == "En Camino" {
					EnviarPaquete(&camion, c, 1)
				}
				if camion.Paquete2.Intentos < 3 && camion.Paquete2.GetEstado() == "En Camino" {
					EnviarPaquete(&camion, c, 2)
				}

				if camion.Paquete1.Intentos < 3 && camion.Paquete1.GetEstado() == "En Camino" {
					EnviarPaquete(&camion, c, 1)
				}
				if camion.Paquete2.Intentos < 3 && camion.Paquete2.GetEstado() == "En Camino" {
					EnviarPaquete(&camion, c, 2)
				}

				////////////////////////////////////
				if camion.Paquete1.Intentos == 3 && camion.Paquete1.GetEstado() == "En Camino" {
					me3 := chat.Message{
						Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
					}
					camion.Paquete1.Estado = "No Recibido"
					log.Printf("%d", camion.Paquete1.Intentos)
					c.CambiarEstado(context.Background(), &me3)
				}
				if camion.Paquete2.Intentos == 3 && camion.Paquete2.GetEstado() == "En Camino" {
					me3 := chat.Message{
						Body: camion.Paquete2.GetTrack() + "%" + "No Recibido",
					}
					camion.Paquete2.Estado = "No Recibido"
					log.Printf("%d", camion.Paquete2.Intentos)
					c.CambiarEstado(context.Background(), &me3)
				}
			} else {
				EnviarPaquete(&camion, c, 2)
				/////////////////////////////////
				EnviarPaquete(&camion, c, 1)

				if camion.Paquete2.Intentos < 3 && camion.Paquete2.GetEstado() == "En Camino" {
					EnviarPaquete(&camion, c, 2)
				}

				if camion.Paquete1.Intentos < 3 && camion.Paquete1.GetEstado() == "En Camino" {
					EnviarPaquete(&camion, c, 1)
				}
				if camion.Paquete2.Intentos < 3 && camion.Paquete2.GetEstado() == "En Camino" {
					EnviarPaquete(&camion, c, 2)
				}

				if camion.Paquete1.Intentos < 3 && camion.Paquete1.GetEstado() == "En Camino" {
					EnviarPaquete(&camion, c, 1)
				}
				////////////////////
				if camion.Paquete1.Intentos == 3 && camion.Paquete1.GetEstado() == "En Camino" {
					me3 := chat.Message{
						Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
					}
					camion.Paquete1.Estado = "No Recibido"
					log.Printf("%d", camion.Paquete1.Intentos)
					c.CambiarEstado(context.Background(), &me3)
				}
				if camion.Paquete2.Intentos == 3 && camion.Paquete2.GetEstado() == "En Camino" {
					me3 := chat.Message{
						Body: camion.Paquete2.GetTrack() + "%" + "No Recibido",
					}
					camion.Paquete2.Estado = "No Recibido"
					log.Printf("%d", camion.Paquete2.Intentos)
					c.CambiarEstado(context.Background(), &me3)
				}

			}
		} else {
			camion.Paquete1 = p1
			log.Printf("Camion %d recibio %d,", camion.id, camion.Paquete1.GetId())
			can.Lock()
			EnviarPaquete(&camion, c, 1)
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
	var i2 int
	var i int

	fmt.Println("Ingrese Tiempo espera Camion")
	_, err := fmt.Scanf("%d\n", &i)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Ingrese Tiempo espera Camion")
	_, err2 := fmt.Scanf("%d\n", &i2)
	if err2 != nil {
		fmt.Println(err2)
	}
	var can sync.Mutex
	for {
		time.Sleep((5) * time.Second)
		can.Lock()
		go Send(CamionRetail1, i, i)
		can.Unlock()
		can.Lock()
		go Send(CamionRetail2, i, i)
		can.Unlock()
	}

}
