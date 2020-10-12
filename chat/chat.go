package chat

import (
	context "context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Server struct {
}

//WriteData is
func WriteData(tipo string, id string, producto string, valor string, inicio string, destino string) {
	t := time.Now()
	timestamp := fmt.Sprintf("%02d-%02d-%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())

	registro := []string{timestamp, id, tipo, producto, valor, inicio, destino}

	csvfile, err := os.OpenFile("dblogistica.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatal(err)
	}

	csvwriter := csv.NewWriter(csvfile)

	csvwriter.Write(registro)

	csvwriter.Flush()
	csvfile.Close()
}

//MandarOrden is
func (s *Server) MandarOrden(ctx context.Context, message *Message) (*Message, error) {
	registro := strings.Split(message.Body, "%")
	id, producto, valor, tienda, destino, tipo := registro[0], registro[1], registro[2], registro[3], registro[4], registro[5]
	log.Printf("Su codigo de tracking es %s", tipo)
	WriteData(tipo, id, producto, valor, tienda, destino)
	trackin := id + "000" + " Para el producto: " + id
	return &Message{Body: trackin}, nil
}

//Consultar is
func (s *Server) Consultar(ctx context.Context, message *Message) (*Message, error) {
	trackin := "000" + " Para el producto: "
	return &Message{Body: trackin}, nil
}
