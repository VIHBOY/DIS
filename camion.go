package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

type Cola struct {
	Nombre   string
	Cola     []string
	Cantidad int
}

type Camion struct {
	Nombre      string
	Paquete1    string
	Paquete2    string
	LlevoRetail string
}

//NewOrden is

func main() {
	ColaRetail := Cola{
		Nombre:   "Retail",
		Cantidad: 0,
	}
	ColaNormal := Cola{
		Nombre:   "Normal",
		Cantidad: 0,
	}
	ColaPrio := Cola{
		Nombre:   "Prioritario",
		Cantidad: 0,
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
	message := chat.Message{
		Body: "%",
	}
	var2 := time.Duration(5) * time.Second
	for {
		time.Sleep(var2)
		response, _ := c.Recibir(context.Background(), &message)
		registro := strings.Split(response.Body, "%")
		id, track, tipo, intentos, estados := registro[0], registro[1], registro[2], registro[3], registro[4]

		time.Sleep(var2)
		response, _ = c.Recibir(context.Background(), &message)
		registro = strings.Split(response.Body, "%")
		id, track, tipo, intentos, estados = registro[0], registro[1], registro[2], registro[3], registro[4]
		if tipo == "retail" {
			log.Printf("Su Recibio IF: %s", tipo)
			append(ColaRetail.Cola, id+"%"+track+"%"+tipo+"%"+intentos+"%"+estados)
			ColaRetail.Cantidad++
		}
		if tipo == "normal" {
			log.Printf("Su Recibio IF: %s", tipo)
			append(ColaNormal.Cola, id+"%"+track+"%"+tipo+"%"+intentos+"%"+estados)
			ColaNormal.Cantidad++
		}
		if tipo == "prioritario" {
			log.Printf("Su Recibio IF: %s", tipo)
			append(ColaPrio.Cola, id+"%"+track+"%"+tipo+"%"+intentos+"%"+estados)
			ColaPrio.Cantidad++
		}
	}

}
