package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

type response2 struct {
	Id       string `json:"id"`
	Track    string `json:"track"`
	Tipo     string `json:"tipo"`
	Valor    int    `json:"valor"`
	Intentos int    `json:"intentos"`
	Estado   string `json:"estado"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
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
			res := response2{}
			str := d.Body
			json.Unmarshal([]byte(str), &res)
			fmt.Println(res)
			log.Printf("Tipo: %s", d.Body)
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

		}

		if strings.Compare("2", text) == 0 {
			break
		}

	}

}
