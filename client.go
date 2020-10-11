package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

type orden struct {
	id       string
	producto string
	valor    string
	tienda   string
	destino  string
}

//NewOrden is
func NewOrden(id string, producto string, valor string, tienda string, destino string) orden {

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

	if err != nil {
		log.Fatalf("no se pudo DECIR HOLA: %s", err)
	}

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

		orden := NewOrden(record[0], record[1], record[2], record[3], record[4])
		fmt.Printf("Holi %s\n", orden.producto)
		message := chat.Message{
			Body: orden.id + "%" + orden.producto + "%" + orden.valor + "%" + orden.tienda + "%" + orden.destino,
		}
		response, err := c.SayHello(context.Background(), &message)
		log.Printf("respuesta del server %s", response.Body)
	}
}
