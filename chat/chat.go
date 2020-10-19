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

//Server is
/***
* struct Server
**
* Estructura Server
**
* Fields
* sync.Mutex mux Herramienta de sincronizacion
* []Paquete ColarNormal2 Lista de Paquetes
* []Paquete ColaPrio2 Lista de Paquetes
* []Paquete ColaRetail2 Lista de Paquetes
* []Orden Lista Lista de Paquetes
* []Paquete ListaTotalCola Lista de Paquetes
**
***/
type Server struct {
	mux            sync.Mutex
	ColaNormal2    []Paquete
	ColaPrio2      []Paquete
	ColaRetail2    []Paquete
	Lista          []Orden
	ListaTotalCola []Paquete
}

/***
* func failOnError
**
* Resumen Devuelve error
**
* Input: err, msg
* error :  error
* string : mensaje
**
***/
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//WriteData is
/***
* func WriteData
**
* Escribe los datos de las ordenes en su csv
**
* string tipo dato de orden
* string id dato de orden
* string producto dato de orden
* string valor dato de orden
* string inicio dato de orden
* string destino dato de orden
***/
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

//Consultar is
func (s *Server) Consultar(ctx context.Context, message *Message) (*Message, error) {
	x := message.GetBody()
	x3 := x
	me := Message{
		Body: x[0 : len(x)-4],
	}

	for i := len(s.ListaTotalCola) - 1; i >= 0; i-- {
		if x3 == s.ListaTotalCola[i].GetTrack() {
			me = Message{
				Body: s.ListaTotalCola[i].GetEstado(),
			}
			break
		}
	}

	return &me, nil
}

//MandarOrden2 is
/***
* func MandarOrden2
**
* Recibe ordenes y crea paquetes
**
* context.Context contexto
* *Orden puntero orden
**
*Devulve un mensaje indicando el track
***/
func (s *Server) MandarOrden2(ctx context.Context, orden *Orden) (*Message, error) {
	s.mux.Lock()
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
	}
	s.mux.Unlock()
	return &me, nil
}

//Recibir2 is
/***
* func Recibir2
**
* Reparte paquetes a los camiones que lo soliciten
**
* context.Context contexto
* *Message puntero message que contiene el tipo de camion
**
*	Devulve un puntero al paquete que sera asignado
***/
func (s *Server) Recibir2(ctx context.Context, message *Message) (*Paquete, error) {
	var me Paquete
	s.mux.Lock()
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
		}

	}

	if message.GetBody() == "Normal" {
		if len(s.ColaPrio2) > 0 {

			me = Paquete{
				Id:       s.ColaPrio2[0].GetId(),
				Track:    s.ColaPrio2[0].GetTrack(),
				Tipo:     s.ColaPrio2[0].GetTipo(),
				Valor:    s.ColaPrio2[0].GetValor(),
				Intentos: s.ColaPrio2[0].GetIntentos(),
				Estado:   s.ColaPrio2[0].GetEstado(),
			}
			if len(s.ColaPrio2) == 1 {
				s.ColaPrio2 = make([]Paquete, 0)
			} else {
				s.ColaPrio2 = s.ColaPrio2[1:]
			}
		} else {
			if len(s.ColaNormal2) > 0 {

				me = Paquete{
					Id:       s.ColaNormal2[0].GetId(),
					Track:    s.ColaNormal2[0].GetTrack(),
					Tipo:     s.ColaNormal2[0].GetTipo(),
					Valor:    s.ColaNormal2[0].GetValor(),
					Intentos: s.ColaNormal2[0].GetIntentos(),
					Estado:   s.ColaNormal2[0].GetEstado(),
				}
				if len(s.ColaNormal2) == 1 {
					s.ColaNormal2 = make([]Paquete, 0)
				} else {
					s.ColaNormal2 = s.ColaNormal2[1:]
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
			}

		}

	}

	if message.GetBody() == "RetailPrio" {
		if len(s.ColaRetail2) > 0 {

			me = Paquete{
				Id:       s.ColaRetail2[0].GetId(),
				Track:    s.ColaRetail2[0].GetTrack(),
				Tipo:     s.ColaRetail2[0].GetTipo(),
				Valor:    s.ColaRetail2[0].GetValor(),
				Intentos: s.ColaRetail2[0].GetIntentos(),
				Estado:   s.ColaRetail2[0].GetEstado(),
			}
			if len(s.ColaRetail2) == 1 {
				s.ColaRetail2 = make([]Paquete, 0)
			} else {
				s.ColaRetail2 = s.ColaRetail2[1:]
			}
		} else {
			if len(s.ColaPrio2) > 0 {

				me = Paquete{
					Id:       s.ColaPrio2[0].GetId(),
					Track:    s.ColaPrio2[0].GetTrack(),
					Tipo:     s.ColaPrio2[0].GetTipo(),
					Valor:    s.ColaPrio2[0].GetValor(),
					Intentos: s.ColaPrio2[0].GetIntentos(),
					Estado:   s.ColaPrio2[0].GetEstado(),
				}
				if len(s.ColaPrio2) == 1 {
					s.ColaPrio2 = make([]Paquete, 0)
				} else {
					s.ColaPrio2 = s.ColaPrio2[1:]
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
			}
		}

	}

	if message.GetBody() == "RetailPrio2" {
		if len(s.ColaPrio2) > 0 {
			me = Paquete{
				Id:       s.ColaPrio2[0].GetId(),
				Track:    s.ColaPrio2[0].GetTrack(),
				Tipo:     s.ColaPrio2[0].GetTipo(),
				Valor:    s.ColaPrio2[0].GetValor(),
				Intentos: s.ColaPrio2[0].GetIntentos(),
				Estado:   s.ColaPrio2[0].GetEstado(),
			}
			if len(s.ColaPrio2) == 1 {
				s.ColaPrio2 = make([]Paquete, 0)
			} else {
				s.ColaPrio2 = s.ColaPrio2[1:]
			}
		} else {
			if len(s.ColaRetail2) > 0 {

				me = Paquete{
					Id:       s.ColaRetail2[0].GetId(),
					Track:    s.ColaRetail2[0].GetTrack(),
					Tipo:     s.ColaRetail2[0].GetTipo(),
					Valor:    s.ColaRetail2[0].GetValor(),
					Intentos: s.ColaRetail2[0].GetIntentos(),
					Estado:   s.ColaRetail2[0].GetEstado(),
				}
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
			}
		}

	}

	s.mux.Unlock()
	return &me, nil
}

//MandarFinanzas is
/***
* func MandarFinanzas
**
* Manda json al servidor Rabbit
**
*
* string pack string con datos del json
**
***/
func MandarFinanzas(pack string) {
	conn, err := amqp.Dial("amqp://admin:admin@dist28:5672/")
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
	failOnError(err, "Failed to publish a message")
}

//CambiarEstado is
/***
* func CambiarEstado
**
* dado un numero de track actualiza los campos
**
*
* devuelve un mensaje vacio
**
***/
func (s *Server) CambiarEstado(ctx context.Context, message *Message) (*Message, error) {
	registro := strings.Split(message.Body, "%")
	//9000%encamino
	track, es := registro[0], registro[1]
	found := 0
	for i := len(s.ListaTotalCola) - 1; i >= 0; i-- {
		if track == s.ListaTotalCola[i].GetTrack() {
			s.ListaTotalCola[i].Estado = es
			if es != "No Recibido" {
				s.ListaTotalCola[i].Intentos++
			}
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

//CambiarIntentos is
/***
* func CambiarIntentos
**
* dado un numero de track actualiza los campos
**
*
* devuelve un mensaje vacio
**
***/
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

//BuscarOrden is
//CambiarIntentos is
/***
* func CambiarIntentos
**
* dado un numero de track busca la orden asociada
**
*
* devuelve un mensaje el cual contiene el inicio y destino de la orden
**
***/
func (s *Server) BuscarOrden(ctx context.Context, message *Message) (*Message, error) {
	x := message.GetBody()
	var me Message
	me = Message{
		Body: "",
	}
	for i := len(s.Lista) - 1; i >= 0; i-- {
		if x == s.Lista[i].GetId() {
			me = Message{
				Body: s.Lista[i].GetInicio() + "%" + s.Lista[i].GetDestino(),
			}
			break
		}
	}

	return &me, nil
}
