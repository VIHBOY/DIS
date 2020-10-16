package chat

import (
	context "context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	sync "sync"
	"time"

	"github.com/streadway/amqp"
)

type Server struct {
	Retail         int
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
	ListaTotalCola []Paquete
	ListaMuRetail  []Paquete2
}

type Paquete2 struct {
	mux      sync.Mutex
	PaqueteR Paquete
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
	//MURO DE BERLINI
	return message, nil
}

//Consultar is
func (s *Server) Consultar(ctx context.Context, message *Message) (*Message, error) {
	x := message.GetBody()
	x3 := x
	log.Println(x)
	me := Message{
		Body: x[0 : len(x)-3],
	}

	for i := len(s.ListaTotalCola) - 1; i >= 0; i-- {
		if x3 == s.ListaTotalCola[i].GetTrack() {
			log.Println(s.ListaTotalCola[i].GetEstado())
		}
	}

	return &me, nil
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
		Inicio:   orden.GetInicio(),
		Destino:  orden.GetDestino(),
		Tipo:     orden.GetTipo(),
	})
	val, _ := strconv.Atoi(orden.GetValor())
	if orden.GetTipo() == "retail" {
		s.ColaRetail2 = append(s.ColaRetail2, Paquete{
			Id:       orden.GetId(),
			Track:    track,
			Tipo:     orden.GetTipo(),
			Valor:    int32(val),
			Intentos: 0,
			Estado:   "En Bodega",
		})
		s.ListaTotalCola = append(s.ListaTotalCola, Paquete{
			Id:       orden.GetId(),
			Track:    track,
			Tipo:     orden.GetTipo(),
			Valor:    int32(val),
			Intentos: 0,
			Estado:   "En Bodega",
		})
		s.CantidadRetail++
	}
	if orden.GetTipo() == "normal" {
		s.ColaNormal2 = append(s.ColaNormal2, Paquete{
			Id:       orden.GetId(),
			Track:    track,
			Tipo:     orden.GetTipo(),
			Valor:    int32(val),
			Intentos: 0,
			Estado:   "En Bodega",
		})
		s.ListaTotalCola = append(s.ListaTotalCola, Paquete{
			Id:       orden.GetId(),
			Track:    track,
			Tipo:     orden.GetTipo(),
			Valor:    int32(val),
			Intentos: 0,
			Estado:   "En Bodega",
		})
		s.CantidadNormal++
	}
	if orden.GetTipo() == "prioritario" {
		s.ColaPrio2 = append(s.ColaPrio2, Paquete{
			Id:       orden.GetId(),
			Track:    track,
			Tipo:     orden.GetTipo(),
			Valor:    int32(val),
			Intentos: 0,
			Estado:   "En Bodega",
		})
		s.ListaTotalCola = append(s.ListaTotalCola, Paquete{
			Id:       orden.GetId(),
			Track:    track,
			Tipo:     orden.GetTipo(),
			Valor:    int32(val),
			Intentos: 0,
			Estado:   "En Bodega",
		})
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

func remove2(slice []Paquete, s int) []Paquete {
	return append(slice[:s], slice[s+1:]...)
}
func (s *Server) Recibir2(ctx context.Context, message *Message) (*Paquete, error) {
	var me Paquete
	if message.GetBody() == "Retail" {
		if len(s.ColaRetail2) > 0 {
			me = Paquete{
				Id:       s.ColaRetail2[0].GetId(),
				Track:    s.ColaRetail2[0].GetTrack(),
				Tipo:     s.ColaRetail2[0].GetTipo(),
				Valor:    s.ColaRetail2[0].GetValor(),
				Intentos: s.ColaRetail2[0].GetIntentos(),
				Estado:   s.ColaRetail2[0].GetEstado(),
			}
<<<<<<< Updated upstream
			if len(s.ColaRetail2) == 1 {
				s.ColaRetail2 = make([]Paquete, 0)
			} else {
				s.ColaRetail2 = s.ColaRetail2[1:]
			}
		} else {
			me = Paquete{
				Id:       "NOHAY",
				Track:    "NOHAY",
				Tipo:     "NOHAY",
				Valor:    0,
				Intentos: 0,
				Estado:   "NOHAY",
			}
=======
			s.ColaRetail2[0]
			me = s.ColaRetail2[0]

			s.Retail++
			s.ColaRetail2 = s.ColaRetail2[1:]
>>>>>>> Stashed changes
		}

	}

	return &me, nil
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

func (s *Server) CambiarEstado(ctx context.Context, message *Message) (*Message, error) {
	registro := strings.Split(message.Body, "%")
	//9000%encamino
	track, es := registro[0], registro[1]
	found := 0
	for i := len(s.ListaTotalCola) - 1; i >= 0; i-- {
		if track == s.ListaTotalCola[i].GetTrack() {
			s.ListaTotalCola[i].Estado = es
			found = i
			break
		}
	}

	if es == "Recibido" {
		MandarFinanzas(fmt.Sprintf(`{"id":"%s", "track":"%s", "tipo":"%s", "valor":%d, "intentos":%d, "estado":"%s"}`, s.ListaTotalCola[found].Id, s.ListaTotalCola[found].Track, s.ListaTotalCola[found].Tipo, s.ListaTotalCola[found].Valor, s.ListaTotalCola[found].Intentos, s.ListaTotalCola[found].Estado))
	}

	if es == "No Recibido" {
		MandarFinanzas(fmt.Sprintf(`{"id":"%s", "track":"%s", "tipo":"%s", "valor":%d, "intentos":%d, "estado":"%s"}`, s.ListaTotalCola[found].Id, s.ListaTotalCola[found].Track, s.ListaTotalCola[found].Tipo, s.ListaTotalCola[found].Valor, s.ListaTotalCola[found].Intentos, s.ListaTotalCola[found].Estado))
	}

	me := Message{
		Body: "",
	}

	return &me, nil
}

func (s *Server) CambiarIntentos(ctx context.Context, message *Message) (*Message, error) {
	registro := strings.Split(message.Body, "%")
	//9000%encamino
	track := registro[0]
	for i := len(s.ListaTotalCola) - 1; i >= 0; i-- {
		if track == s.ListaTotalCola[i].GetTrack() {
			s.ListaTotalCola[i].Intentos++
			break
		}
	}

	me := Message{
		Body: "",
	}

	return &me, nil
}
