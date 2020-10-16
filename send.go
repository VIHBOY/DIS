package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

type response2 struct {
	id       string
	track    string
	tipo     string
	valor    string
	intentos int
	estado   string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func MandarFinanzas(pack string) {
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

	body := pack
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	Id := "1"
	Track := "1000"
	Tipo := "Polera"
	Valor := rand.Intn(6-1) + 1
	Intentos := 3
	Estado := "Recibido"
	MandarFinanzas(fmt.Sprintf(`{"id":"%s", "track":"%s", "tipo":"%s", "valor":%d, "intentos":%d, "estado":"%s"}`, Id, Track, Tipo, Valor, Intentos, Estado))
}
