package chat

import (
	context "context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

type Server struct {
	Lista1         []string
	ColaNormal     []string
	ColaNormal2    []Paquete
	ColaPrio2      []Paquete
	ColaRetail2    []Paquete
	CantidadNormal int
	ColaPrio       []string
	CantidadPrio   int
	ColaRetail     []string
	CantidadRetail int
	Lista          []Orden
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
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
	id, producto, valor, inicio, destino, tipo := registro[0], registro[1], registro[2], registro[3], registro[4], registro[5]
	log.Printf("Su codigo de tracking es %s", id)
	WriteData(tipo, id, producto, valor, inicio, destino)
	trackin := id + "000" + " Para el producto: " + id
	//MURO DE BERLINI
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

	body := "{name:max, message:win}"
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
	return &Message{Body: trackin}, nil
}

//Consultar is
func (s *Server) Consultar(ctx context.Context, message *Message) (*Message, error) {
	trackin := "000" + " Para el producto: "
	return &Message{Body: trackin}, nil
}

//MandarOrden2 is
func (s *Server) MandarOrden2(ctx context.Context, orden *Orden) (*Message, error) {
	trackin := orden.GetId() + "000" + " Para el producto: " + orden.GetId()
	me := Message{
		Body: trackin,
	}
	track := orden.GetId() + "000"
	WriteData(orden.GetTipo(), orden.GetId(), orden.GetProducto(), orden.GetValor(), orden.GetInicio(), orden.GetDestino())
	s.Lista = append(s.Lista, Orden{
		Id:       orden.GetId(),
		Producto: orden.GetProducto(),
		Valor:    orden.GetValor(),
		Inicio:   orden.GetValor(),
		Destino:  orden.GetDestino(),
		Tipo:     orden.GetTipo(),
	})
	Body := orden.GetId() + "%" + track + "%" + orden.GetTipo() + "%" + "0" + "%" + "En Bodega"
	paquete := Paquete{
		Id:       orden.GetId(),
		Track:    track,
		Tipo:     orden.GetTipo(),
		Intentos: 0,
		Estado:   "En Bodega",
	}
	fmt.Printf(Paquete.GetId())
	if orden.GetTipo() == "retail" {
		s.ColaRetail = append(s.ColaRetail, Body)
		s.ColaRetail2 = append(s.Lista, paquete)
		s.CantidadRetail++
	}
	if orden.GetTipo() == "normal" {
		s.ColaNormal = append(s.ColaNormal, Body)
		s.CantidadNormal++
	}
	if orden.GetTipo() == "prioritario" {
		s.ColaPrio = append(s.ColaPrio, Body)
		s.CantidadPrio++
	}
	return &me, nil
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
func (s *Server) Recibir(ctx context.Context, message *Message) (*Message, error) {
	me := Message{
		Body: "",
	}
	if message.GetBody() == "Retail" {
		if s.CantidadRetail > 0 {
			me = Message{
				Body: s.ColaRetail[0],
			}
			remove(s.ColaRetail, 0)
			s.CantidadRetail--
		}

	}
	if message.GetBody() == "Normal" {
		if s.CantidadPrio > 0 {
			me = Message{
				Body: s.ColaPrio[0],
			}
			remove(s.ColaPrio, 0)
			s.CantidadPrio--
		} else {
			me = Message{
				Body: s.ColaRetail[0],
			}
			s.CantidadRetail--
			remove(s.ColaRetail, 0)
		}

	}
	return &me, nil
}
