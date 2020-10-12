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
	timestamp := fmt.Sprintf("%02d-%02d-%d %02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())

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

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	registro := strings.Split(message.Body, "%")
	id, producto, valor, inicio, destino := registro[0], registro[1], registro[2], registro[3], registro[4]
	log.Printf("Rece mensaje form client: %s", message.Body)
	log.Printf("%s %s %s %s %s", id, producto, valor, inicio, destino)
	WriteData("retail", id, producto, valor, inicio, destino)
	trackin := id + "000"
	return &Message{Body: trackin}, nil
}
