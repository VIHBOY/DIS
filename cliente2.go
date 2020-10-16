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
	tipo     string
}

//NewOrden is
func NewOrden(id string, producto string, valor string, tienda string, destino string, tipo string) orden {

	o := orden{id: id, producto: producto, valor: valor, tienda: tienda, destino: destino, tipo: tipo}
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
		fmt.Print("3. Consultar \n")

		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)
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

				message := chat.Orden{
					Id:       record[0],
					Producto: record[1],
					Valor:    record[2],
					Inicio:   record[3],
					Destino:  record[4],
					Tipo:     "retail",
				}
				response, err := c.MandarOrden2(context.Background(), &message)
				log.Printf("Su codigo de tracking es %s", response.Body)
			}
		}
		if strings.Compare("2", text) == 0 {
			csvfile, _ := os.Open("pymes.csv")
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
				tipo := ""
				if record[5] == "0" {
					tipo = "normal"
				}
				if record[5] == "1" {
					tipo = "prioritario"
				}
				orden := NewOrden(record[0], record[1], record[2], record[3], record[4], tipo)
				message := chat.Message{
					Body: orden.id + "%" + orden.producto + "%" + orden.valor + "%" + orden.tienda + "%" + orden.destino + "%" + orden.tipo,
				}
				response, err := c.MandarOrden(context.Background(), &message)
				log.Printf("Su codigo de tracking es %s", response.Body)
			}
		}
		if strings.Compare("3", text) == 0 {
			text2, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text2 = strings.Replace(text2, "\r\n", "", -1)
			message := chat.Message{
				Body: text2,
			}
			response, _ := c.Consultar(context.Background(), &message)
			log.Printf("El estado de su producto es: %s", response.Body)
		}

		if strings.Compare("exit", text) == 0 {
			break
		}

	}

}
