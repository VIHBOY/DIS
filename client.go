package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Menú")
		fmt.Println("---------------------")
		fmt.Print("1. Cargar Retail \n")
		fmt.Print("2. Cargar Pymes \n")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("1", text) == 0 {
			csvfile, _ := os.Open("retail.csv")
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
				message := chat.Message{
					Body: orden.id + "%" + orden.producto + "%" + orden.valor + "%" + orden.tienda + "%" + orden.destino,
				}
				response, err := c.SayHello(context.Background(), &message)
				log.Printf("Su codigo de tracking %s", response.Body)
			}
		}

		if strings.Compare("2", text) == 0 {
			fmt.Println("hello, Yourself")
		}

		if strings.Compare("exit", text) == 0 {
			break
		}

	}

}
