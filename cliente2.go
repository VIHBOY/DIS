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
	"time"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %s", err)
	}
	defer conn.Close()
	var i2 int
	c := chat.NewChatServiceClient(conn)

	if err != nil {
		log.Fatalf("no se pudo DECIR HOLA: %s", err)
	}
	var i3 int
	fmt.Println("Ingrese Tipo de Cliente")
	fmt.Println("1. Retail")
	fmt.Println("2. Pyme")
	_, err3 := fmt.Scanf("%d\n", &i3)
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println("Ingrese Tiempo envio de ordenes")
	_, err2 := fmt.Scanf("%d\n", &i2)
	if err2 != nil {
		fmt.Println(err2)
	}
	var tipo string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Menú")
		fmt.Println("---------------------")
		fmt.Print("1. Cargar Datos  \n")
		fmt.Print("2. Consultar \n")
		fmt.Print("3. Salir \n")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)
		if strings.Compare("1", text) == 0 {
			var arch string
			if i3 == 1 {
				arch = "retail.csv"
			}
			if i3 == 2 {
				arch = "pymes.csv"

			}
			csvfile, _ := os.Open(arch)
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
				if i3 == 2 {
					if record[5] == "0" {
						tipo = "normal"
					} else if record[5] == "1" {
						tipo = "prioritario"
					}
				}
				if i3 == 1 {
					tipo = "retail"
				}

				message := chat.Orden{
					Id:       record[0],
					Producto: record[1],
					Valor:    record[2],
					Inicio:   record[3],
					Destino:  record[4],
					Tipo:     tipo,
				}
				response, err := c.MandarOrden2(context.Background(), &message)
				log.Printf("Su codigo de tracking es %s", response.Body)
				time.Sleep(time.Duration(i2) * time.Second)
			}
		}

		if strings.Compare("2", text) == 0 {
			text2, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text2 = strings.Replace(text2, "\r\n", "", -1)
			message := chat.Message{
				Body: text2,
			}
			response, _ := c.Consultar(context.Background(), &message)
			log.Printf("El estado de su producto es: %s", response.Body)
		}

		if strings.Compare("3", text) == 0 {
			break
		}

	}

}
