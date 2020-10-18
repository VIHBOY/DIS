package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
)

//PackJSON is a struct.
type PackJSON struct {
	Id       string  `json:"id"`
	Track    string  `json:"track"`
	Tipo     string  `json:"tipo"`
	Valor    int     `json:"valor"`
	Intentos int     `json:"intentos"`
	Estado   string  `json:"estado"`
	Ganancia float64 `json:"ganancia"`
	Perdida  float64 `json:"perdida"`
	Total    float64 `json:"total"`
}

//WriteData is
func WriteData(name string, id string, track string, tipo string, valor int, intentos int, estado string, ganancia float64, perdida float64, total float64) {
	registro := []string{id, track, tipo, strconv.Itoa(valor), strconv.Itoa(intentos), estado, strconv.FormatFloat(ganancia, 'f', 1, 64), strconv.FormatFloat(perdida, 'f', 1, 64), strconv.FormatFloat(total, 'f', 1, 64)}
	csvfile, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatal(err)
	}

	csvwriter := csv.NewWriter(csvfile)

	csvwriter.Write(registro)

	csvwriter.Flush()
	csvfile.Close()
}

//CreateFile is
func CreateFile(name string) {

	csvFile, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	csvFile.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	CreateFile("finanzas.csv")
	ListaPJ := []PackJSON{}
	var tGanancia float64
	var tPerdida float64
	var tTotal float64
	var LenListaPJ int
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Finanzas", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			res := PackJSON{}
			str := d.Body
			json.Unmarshal([]byte(str), &res)
			if res.Estado == "Recibido" {
				if res.Tipo == "retail" {
					ListaPJ = append(ListaPJ, PackJSON{
						Id:       res.Id,
						Track:    res.Track,
						Tipo:     res.Tipo,
						Valor:    res.Valor,
						Intentos: res.Intentos,
						Estado:   res.Estado,
						Ganancia: float64(res.Valor),
						Perdida:  float64((res.Intentos - 1) * 10),
						Total:    float64(res.Valor - (res.Intentos-1)*10),
					})
				} else if res.Tipo == "normal" {
					ListaPJ = append(ListaPJ, PackJSON{
						Id:       res.Id,
						Track:    res.Track,
						Tipo:     res.Tipo,
						Valor:    res.Valor,
						Intentos: res.Intentos,
						Estado:   res.Estado,
						Ganancia: float64(res.Valor),
						Perdida:  float64((res.Intentos - 1) * 10),
						Total:    float64(res.Valor - (res.Intentos-1)*10),
					})
				} else {
					ListaPJ = append(ListaPJ, PackJSON{
						Id:       res.Id,
						Track:    res.Track,
						Tipo:     res.Tipo,
						Valor:    res.Valor,
						Intentos: res.Intentos,
						Estado:   res.Estado,
						Ganancia: float64(res.Valor) * 1.3,
						Perdida:  float64((res.Intentos - 1) * 10),
						Total:    float64(float64(res.Valor)*1.3 - float64((res.Intentos-1)*10)),
					})
				}
			} else {
				if res.Tipo == "retail" {
					ListaPJ = append(ListaPJ, PackJSON{
						Id:       res.Id,
						Track:    res.Track,
						Tipo:     res.Tipo,
						Valor:    res.Valor,
						Intentos: res.Intentos,
						Estado:   res.Estado,
						Ganancia: float64(res.Valor),
						Perdida:  float64(0),
						Total:    float64(res.Valor),
					})

				} else if res.Tipo == "normal" {
					ListaPJ = append(ListaPJ, PackJSON{
						Id:       res.Id,
						Track:    res.Track,
						Tipo:     res.Tipo,
						Valor:    res.Valor,
						Intentos: res.Intentos,
						Estado:   res.Estado,
						Ganancia: float64(0),
						Perdida:  float64((res.Intentos - 1) * 10),
						Total:    float64(-(res.Intentos - 1) * 10),
					})
				} else {
					ListaPJ = append(ListaPJ, PackJSON{
						Id:       res.Id,
						Track:    res.Track,
						Tipo:     res.Tipo,
						Valor:    res.Valor,
						Intentos: res.Intentos,
						Estado:   res.Estado,
						Ganancia: float64(res.Valor) * 0.3,
						Perdida:  float64((res.Intentos - 1) * 10),
						Total:    float64(float64(res.Valor)*0.3 - float64((res.Intentos-1)*10)),
					})
				}
			}
			LenListaPJ = len(ListaPJ)
			tGanancia += ListaPJ[LenListaPJ-1].Ganancia
			tPerdida += ListaPJ[LenListaPJ-1].Perdida
			tTotal += ListaPJ[LenListaPJ-1].Total
			WriteData("finanzas.csv", ListaPJ[LenListaPJ-1].Id, ListaPJ[LenListaPJ-1].Track, ListaPJ[LenListaPJ-1].Tipo, ListaPJ[LenListaPJ-1].Valor, ListaPJ[LenListaPJ-1].Intentos, ListaPJ[LenListaPJ-1].Estado, ListaPJ[LenListaPJ-1].Ganancia, ListaPJ[LenListaPJ-1].Perdida, ListaPJ[LenListaPJ-1].Total)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Menú")
		fmt.Println("---------------------")
		fmt.Print("1. Consultar estado \n")
		fmt.Print("2. Salir \n")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)
		if strings.Compare("1", text) == 0 {
			fmt.Println(ListaPJ)
			fmt.Println(tGanancia, tPerdida, tTotal)
		}

		if strings.Compare("2", text) == 0 {
			break
		}

	}

}
