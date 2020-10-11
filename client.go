package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

type orden struct {
	id       string
	producto string
	valor    int
	tienda   string
	destino  string
}

func NewOrden(id string, producto string, valor int, tienda string, destino string) orden {

	o := orden{id: id, producto: producto, valor: valor, tienda: tienda, destino: destino}
	return o
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Holi soy el cliente",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("no se pudo DECIR HOLA: %s", err)
	}
	log.Printf("respuesta del server %s", response.Body)

	csvfile, err := os.Open("retail.csv")
	r := csv.NewReader(csvfile)
	r.Read()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		val, err := strconv.Atoi(record[2])

		orden := NewOrden(record[0], record[1], val, record[3], record[4])
		fmt.Printf("Holi %s\n", orden.producto)
		message := chat.Message{
			Body: "Holi soy el "+orden.producto,
		}
		response, err := c.SayHello(context.Background(), &message)
		log.Printf("respuesta del server %s", response.Body)
	}
}
