package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/VIHBOY/DIS/chat"
	"google.golang.org/grpc"
)

//Camion is
/***
* struct Camion
**
* Estructura de camiones
**
* Fields:
* sync.Mutex : Herramienta para sincronizacion
* int id : Id del camion
* string Nombre : Tipo del camion
* *chat.Paquete Paquete1 : Puntero al primer paquete
* *chat.Paquete Paquete2 : Puntero al segundo paquete
* int LleveRetail : Flag que indica si se envio paquete retail en el envio anterior
* int Deruta : Flag que indica si el camion esta ocupado
* string NombreArchivo : Nombre del archivo CSV del camion
***/
type Camion struct {
	mux           sync.Mutex
	id            int
	Nombre        string
	Paquete1      *chat.Paquete
	Paquete2      *chat.Paquete
	LleveRetail   int
	Deruta        int
	NombreArchivo string
}

//WriteData2 is
/***
* func WriteData2
**
* Escribe datos de camiones en archivos
**
* Input:
* string name : Nombre del archivo
* *chat.Paquete paquete : Puntero al paquete
* *(chat.Message) auxiliar : Puntero de un response de logistica
***/
func WriteData2(name string, paquete *chat.Paquete, auxiliar *(chat.Message)) {
	registro2 := strings.Split(auxiliar.Body, "%")
	inicio, destino := registro2[0], registro2[1]
	t := time.Now()
	timestamp := fmt.Sprintf("%02d-%02d-%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	csvfile, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	registro := []string{timestamp, paquete.GetId(), paquete.GetTipo(), strconv.Itoa(int(paquete.GetValor())), inicio, destino, timestamp}
	if err != nil {
		log.Fatal(err)
	}

	csvwriter := csv.NewWriter(csvfile)

	csvwriter.Write(registro)

	csvwriter.Flush()
	csvfile.Close()
}

//CreateFile is
/***
* func CreateFile
**
* Crea archivos APPEND
**
* Input:
* string name : Nombre del archivo CSV
***/
func CreateFile(name string) {

	csvFile, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	csvFile.Close()
}

//EnviarPaquete is
/***
* func EnviarPaquete
**
* Envia paquetes a destino
**
* Input:
* *Camion camion : Puntero a un camion
* chat.ChatServiceClient c : Llamada a funciones de chat.go
* int np : Numero del paquete (1 o 2)
***/
func EnviarPaquete(camion *Camion, c chat.ChatServiceClient, np int) {
	rand.Seed(time.Now().UnixNano())
	prob := rand.Intn(6-1) + 1

	if np == 1 {
		if prob == 5 {
			me3 := chat.Message{
				Body: camion.Paquete1.GetTrack() + "%" + "En Camino",
			}
			camion.Paquete1.Estado = "En Camino"
			camion.Paquete1.Intentos++
			c.CambiarEstado(context.Background(), &me3)
			log.Printf("Camion: %d: Paquete: %s Intentos %d:\n", camion.id, camion.Paquete1.Id, camion.Paquete1.Intentos)

		} else {
			me3 := chat.Message{
				Body: camion.Paquete1.GetTrack() + "%" + "Recibido",
			}
			camion.Paquete1.Estado = "Recibido"
			camion.Paquete1.Intentos++
			c.CambiarEstado(context.Background(), &me3)
			log.Printf("Camion: %d: Paquete: %s Intentos %d:\n", camion.id, camion.Paquete1.Id, camion.Paquete1.Intentos)

		}
	}
	if np == 2 {
		if prob == 5 {
			me3 := chat.Message{
				Body: camion.Paquete2.GetTrack() + "%" + "En Camino",
			}
			camion.Paquete2.Estado = "En Camino"
			camion.Paquete2.Intentos++
			c.CambiarEstado(context.Background(), &me3)
			log.Printf("Camion: %d: Paquete: %s Intentos %d:\n", camion.id, camion.Paquete2.Id, camion.Paquete2.Intentos)

		} else {
			me3 := chat.Message{
				Body: camion.Paquete2.GetTrack() + "%" + "Recibido",
			}
			camion.Paquete2.Estado = "Recibido"
			camion.Paquete2.Intentos++
			c.CambiarEstado(context.Background(), &me3)
			log.Printf("Camion: %d: Paquete: %s Intentos %d:\n", camion.id, camion.Paquete2.Id, camion.Paquete2.Intentos)

		}
	}

}

//Send is
/***
* func Send
**
* Envia paquetes y le informa el resultado a logistica
**
* Input:
* *Camion camion : Puntero a un camion
* int tespera : Tiempo de espera por un segundo paquete del camion
* int tenvio : Tiempo que demora en entregar un paquete un camion
***/
func Send(camion *Camion, tespera int, tenvio int) {
	var conn *grpc.ClientConn
	var me chat.Message
	conn, _ = grpc.Dial("dist25:9000", grpc.WithInsecure())
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: camion.Nombre,
	}

	if camion.id == 1 || camion.id == 2 {
		if camion.LleveRetail == 1 {
			message = chat.Message{
				Body: "RetailPrio",
			}
		}
	}
	p1, _ := c.Recibir2(context.Background(), &message)

	if camion.id == 1 || camion.id == 2 {
		if camion.LleveRetail == 1 {

			message = chat.Message{
				Body: "RetailPrio2",
			}
		}
	}
	p2, _ := c.Recibir2(context.Background(), &message)

	if p1.GetId() != "NOHAY" {
		if p2.GetId() != "NOHAY" {
			camion.Paquete1 = p1
			camion.Paquete2 = p2
			log.Printf("Camion %d 1 recibio %s,", camion.id, camion.Paquete1.GetId())
			log.Printf("Camion %d 2 recibio %s,", camion.id, camion.Paquete2.GetId())
			if camion.Paquete1.GetValor() >= camion.Paquete2.GetValor() {
				time.Sleep(time.Duration(tenvio) * time.Second)
				EnviarPaquete(camion, c, 1)
				/////////////////////////////////
				time.Sleep(time.Duration(tenvio) * time.Second)
				EnviarPaquete(camion, c, 2)

				if camion.Paquete1.Intentos < 3 && camion.Paquete1.GetEstado() == "En Camino" {

					if camion.Paquete1.GetTipo() == "normal" || camion.Paquete1.GetTipo() == "prioritario" {
						log.Printf("Valor %d Paquete: %d", camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10, camion.Paquete1.Id)
						if camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10 >= 0 {
							time.Sleep(time.Duration(tenvio) * time.Second)
							EnviarPaquete(camion, c, 1)
						} else {
							me3 := chat.Message{
								Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
							}
							camion.Paquete1.Estado = "No Recibido"
							log.Printf("%d", camion.Paquete1.Intentos)
							c.CambiarEstado(context.Background(), &me3)
						}

					} else {
						time.Sleep(time.Duration(tenvio) * time.Second)
						EnviarPaquete(camion, c, 1)
					}

				}
				if camion.Paquete2.Intentos < 3 && camion.Paquete2.GetEstado() == "En Camino" {

					if camion.Paquete2.GetTipo() == "normal" || camion.Paquete2.GetTipo() == "prioritario" {
						log.Printf("Valor %d Paquete: %d", camion.Paquete2.GetValor()-(camion.Paquete2.GetIntentos())*10, camion.Paquete2.Id)
						if camion.Paquete2.GetValor()-(camion.Paquete2.GetIntentos())*10 >= 0 {
							time.Sleep(time.Duration(tenvio) * time.Second)
							EnviarPaquete(camion, c, 2)
						} else {
							me3 := chat.Message{
								Body: camion.Paquete2.GetTrack() + "%" + "No Recibido",
							}
							camion.Paquete2.Estado = "No Recibido"
							log.Printf("%d", camion.Paquete2.Intentos)
							c.CambiarEstado(context.Background(), &me3)
						}

					} else {
						time.Sleep(time.Duration(tenvio) * time.Second)
						EnviarPaquete(camion, c, 2)
					}

				}

				if camion.Paquete1.Intentos < 3 && camion.Paquete1.GetEstado() == "En Camino" {

					if camion.Paquete1.GetTipo() == "normal" || camion.Paquete1.GetTipo() == "prioritario" {
						log.Printf("Valor %d Paquete: %d", camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10, camion.Paquete1.Id)
						if camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10 >= 0 {
							time.Sleep(time.Duration(tenvio) * time.Second)
							EnviarPaquete(camion, c, 1)
						} else {
							me3 := chat.Message{
								Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
							}
							camion.Paquete1.Estado = "No Recibido"
							log.Printf("%d", camion.Paquete1.Intentos)
							c.CambiarEstado(context.Background(), &me3)
						}

					} else {
						time.Sleep(time.Duration(tenvio) * time.Second)
						EnviarPaquete(camion, c, 1)
					}
				}
				if camion.Paquete2.Intentos < 3 && camion.Paquete2.GetEstado() == "En Camino" {

					if camion.Paquete2.GetTipo() == "normal" || camion.Paquete2.GetTipo() == "prioritario" {
						log.Printf("Valor %d Paquete: %d", camion.Paquete2.GetValor()-(camion.Paquete2.GetIntentos())*10, camion.Paquete2.Id)
						if camion.Paquete2.GetValor()-(camion.Paquete2.GetIntentos())*10 >= 0 {
							time.Sleep(time.Duration(tenvio) * time.Second)
							EnviarPaquete(camion, c, 2)
						} else {
							me3 := chat.Message{
								Body: camion.Paquete2.GetTrack() + "%" + "No Recibido",
							}
							camion.Paquete2.Estado = "No Recibido"
							log.Printf("%d", camion.Paquete2.Intentos)
							c.CambiarEstado(context.Background(), &me3)
						}

					} else {
						time.Sleep(time.Duration(tenvio) * time.Second)
						EnviarPaquete(camion, c, 2)
					}

				}

				////////////////////////////////////
				if camion.Paquete1.Intentos == 3 && camion.Paquete1.GetEstado() == "En Camino" {
					me3 := chat.Message{
						Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
					}
					camion.Paquete1.Estado = "No Recibido"
					log.Printf("%d", camion.Paquete1.Intentos)
					c.CambiarEstado(context.Background(), &me3)
				}
				if camion.Paquete2.Intentos == 3 && camion.Paquete2.GetEstado() == "En Camino" {
					me3 := chat.Message{
						Body: camion.Paquete2.GetTrack() + "%" + "No Recibido",
					}
					camion.Paquete2.Estado = "No Recibido"
					log.Printf("%d", camion.Paquete2.Intentos)
					c.CambiarEstado(context.Background(), &me3)
				}
				camion.Deruta = 0
				if camion.Paquete1.GetTipo() == "retail" || camion.Paquete2.GetTipo() == "retail" {
					camion.LleveRetail = 1
				}
				if camion.Paquete1.GetTipo() == "prioritario" && camion.Paquete2.GetTipo() == "prioritario" {
					camion.LleveRetail = 0
				}
				me = chat.Message{
					Body: camion.Paquete1.GetId(),
				}
				respuesta, _ := c.BuscarOrden(context.Background(), &me)
				WriteData2(camion.NombreArchivo, camion.Paquete1, respuesta)
				me = chat.Message{
					Body: camion.Paquete2.GetId(),
				}
				respuesta, _ = c.BuscarOrden(context.Background(), &me)
				WriteData2(camion.NombreArchivo, camion.Paquete2, respuesta)
			} else {
				time.Sleep(time.Duration(tenvio) * time.Second)
				EnviarPaquete(camion, c, 2)
				/////////////////////////////////
				time.Sleep(time.Duration(tenvio) * time.Second)

				EnviarPaquete(camion, c, 1)

				if camion.Paquete2.Intentos < 3 && camion.Paquete2.GetEstado() == "En Camino" {

					if camion.Paquete2.GetTipo() == "normal" || camion.Paquete2.GetTipo() == "prioritario" {
						log.Printf("Valor %d Paquete: %d", camion.Paquete2.GetValor()-(camion.Paquete2.GetIntentos())*10, camion.Paquete2.Id)
						if camion.Paquete2.GetValor()-(camion.Paquete2.GetIntentos())*10 >= 0 {
							time.Sleep(time.Duration(tenvio) * time.Second)
							EnviarPaquete(camion, c, 2)
						} else {
							me3 := chat.Message{
								Body: camion.Paquete2.GetTrack() + "%" + "No Recibido",
							}
							camion.Paquete2.Estado = "No Recibido"
							log.Printf("%d", camion.Paquete2.Intentos)
							c.CambiarEstado(context.Background(), &me3)
						}

					} else {
						time.Sleep(time.Duration(tenvio) * time.Second)
						EnviarPaquete(camion, c, 2)
					}

				}

				if camion.Paquete1.Intentos < 3 && camion.Paquete1.GetEstado() == "En Camino" {

					if camion.Paquete1.GetTipo() == "normal" || camion.Paquete1.GetTipo() == "prioritario" {
						log.Printf("Valor %d Paquete: %d", camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10, camion.Paquete1.Id)
						if camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10 >= 0 {
							time.Sleep(time.Duration(tenvio) * time.Second)
							EnviarPaquete(camion, c, 1)
						} else {
							me3 := chat.Message{
								Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
							}
							camion.Paquete1.Estado = "No Recibido"
							log.Printf("%d", camion.Paquete1.Intentos)
							c.CambiarEstado(context.Background(), &me3)
						}

					} else {
						time.Sleep(time.Duration(tenvio) * time.Second)
						EnviarPaquete(camion, c, 1)
					}
				}
				if camion.Paquete2.Intentos < 3 && camion.Paquete2.GetEstado() == "En Camino" {

					if camion.Paquete2.GetTipo() == "normal" || camion.Paquete2.GetTipo() == "prioritario" {
						log.Printf("Valor %d Paquete: %d", camion.Paquete2.GetValor()-(camion.Paquete2.GetIntentos())*10, camion.Paquete2.Id)
						if camion.Paquete2.GetValor()-(camion.Paquete2.GetIntentos())*10 >= 0 {
							time.Sleep(time.Duration(tenvio) * time.Second)
							EnviarPaquete(camion, c, 2)
						} else {
							me3 := chat.Message{
								Body: camion.Paquete2.GetTrack() + "%" + "No Recibido",
							}
							camion.Paquete2.Estado = "No Recibido"
							log.Printf("%d", camion.Paquete2.Intentos)
							c.CambiarEstado(context.Background(), &me3)
						}

					} else {
						time.Sleep(time.Duration(tenvio) * time.Second)
						EnviarPaquete(camion, c, 2)
					}

				}

				if camion.Paquete1.Intentos < 3 && camion.Paquete1.GetEstado() == "En Camino" {
					if camion.Paquete1.GetTipo() == "normal" || camion.Paquete1.GetTipo() == "prioritario" {
						log.Printf("Valor %d Paquete: %d", camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10, camion.Paquete1.Id)
						if camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10 >= 0 {
							time.Sleep(time.Duration(tenvio) * time.Second)
							EnviarPaquete(camion, c, 1)
						} else {
							me3 := chat.Message{
								Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
							}
							camion.Paquete1.Estado = "No Recibido"
							log.Printf("%d", camion.Paquete1.Intentos)
							c.CambiarEstado(context.Background(), &me3)
						}

					} else {
						time.Sleep(time.Duration(tenvio) * time.Second)
						EnviarPaquete(camion, c, 1)
					}
				}
				////////////////////
				if camion.Paquete1.Intentos == 3 && camion.Paquete1.GetEstado() == "En Camino" {
					me3 := chat.Message{
						Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
					}
					camion.Paquete1.Estado = "No Recibido"
					log.Printf("%d", camion.Paquete1.Intentos)
					c.CambiarEstado(context.Background(), &me3)
				}
				if camion.Paquete2.Intentos == 3 && camion.Paquete2.GetEstado() == "En Camino" {
					me3 := chat.Message{
						Body: camion.Paquete2.GetTrack() + "%" + "No Recibido",
					}
					camion.Paquete2.Estado = "No Recibido"
					log.Printf("%d", camion.Paquete2.Intentos)
					c.CambiarEstado(context.Background(), &me3)
				}
				camion.Deruta = 0
				if camion.Paquete1.GetTipo() == "retail" || camion.Paquete2.GetTipo() == "retail" {
					camion.LleveRetail = 1
				}
				if camion.Paquete1.GetTipo() == "prioritario" && camion.Paquete2.GetTipo() == "prioritario" {
					camion.LleveRetail = 0
				}
				me = chat.Message{
					Body: camion.Paquete1.GetId(),
				}
				respuesta, _ := c.BuscarOrden(context.Background(), &me)
				WriteData2(camion.NombreArchivo, camion.Paquete1, respuesta)
				me = chat.Message{
					Body: camion.Paquete2.GetId(),
				}
				respuesta, _ = c.BuscarOrden(context.Background(), &me)
				WriteData2(camion.NombreArchivo, camion.Paquete2, respuesta)
			}
		} else {
			camion.Paquete1 = p1
			log.Printf("Camion %d recibio %s,", camion.id, camion.Paquete1.GetId())
			time.Sleep(time.Duration(tenvio) * time.Second)

			EnviarPaquete(camion, c, 1)
			if camion.Paquete1.Intentos < 3 && camion.Paquete1.GetEstado() == "En Camino" {

				if camion.Paquete1.GetTipo() == "normal" || camion.Paquete1.GetTipo() == "prioritario" {
					log.Printf("Valor %d Paquete: %d", camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10, camion.Paquete1.Id)
					if camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10 >= 0 {
						time.Sleep(time.Duration(tenvio) * time.Second)
						EnviarPaquete(camion, c, 1)
					} else {
						me3 := chat.Message{
							Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
						}
						camion.Paquete1.Estado = "No Recibido"
						log.Printf("%d", camion.Paquete1.Intentos)
						c.CambiarEstado(context.Background(), &me3)
					}

				} else {
					time.Sleep(time.Duration(tenvio) * time.Second)
					EnviarPaquete(camion, c, 1)
				}
			}
			if camion.Paquete1.Intentos < 3 && camion.Paquete1.GetEstado() == "En Camino" {

				if camion.Paquete1.GetTipo() == "normal" || camion.Paquete1.GetTipo() == "prioritario" {
					log.Printf("Valor %d Paquete: %d", camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10, camion.Paquete1.Id)
					if camion.Paquete1.GetValor()-(camion.Paquete1.GetIntentos())*10 >= 0 {
						time.Sleep(time.Duration(tenvio) * time.Second)
						EnviarPaquete(camion, c, 1)
					} else {
						me3 := chat.Message{
							Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
						}
						camion.Paquete1.Estado = "No Recibido"
						log.Printf("%d", camion.Paquete1.Intentos)
						c.CambiarEstado(context.Background(), &me3)
					}

				} else {
					time.Sleep(time.Duration(tenvio) * time.Second)
					EnviarPaquete(camion, c, 1)
				}
			}
			if camion.Paquete1.Intentos == 3 && camion.Paquete1.GetEstado() == "En Camino" {
				me3 := chat.Message{
					Body: camion.Paquete1.GetTrack() + "%" + "No Recibido",
				}
				camion.Paquete1.Estado = "No Recibido"
				log.Printf("%d", camion.Paquete1.Intentos)
				c.CambiarEstado(context.Background(), &me3)
			}
			camion.Deruta = 0
			if camion.Paquete1.GetTipo() == "retail" {
				camion.LleveRetail = 1
			}
			if camion.Paquete1.GetTipo() == "prioritario" {
				camion.LleveRetail = 0
			}
			me = chat.Message{
				Body: camion.Paquete1.GetId(),
			}
			respuesta, _ := c.BuscarOrden(context.Background(), &me)
			WriteData2(camion.NombreArchivo, camion.Paquete1, respuesta)

		}
	} else {
		camion.Deruta = 0

	}
}

func main() {
	CreateFile("CamionRetail1.csv")
	CreateFile("CamionRetail2.csv")
	CreateFile("CamionNormal.csv")
	CamionRetail1 := Camion{
		id:            1,
		Nombre:        "Retail",
		Deruta:        0,
		LleveRetail:   0,
		NombreArchivo: "CamionRetail1.csv",
	}
	CamionRetail2 := Camion{
		id:            2,
		Nombre:        "Retail",
		Deruta:        0,
		LleveRetail:   0,
		NombreArchivo: "CamionRetail2.csv",
	}
	CamionNormal := Camion{
		id:            3,
		Nombre:        "Normal",
		Deruta:        0,
		NombreArchivo: "CamionNormal.csv",
	}
	var i2 int
	var i int

	fmt.Println("Ingrese Tiempo de espera por un segundo paquete del camion")
	_, err := fmt.Scanf("%d\n", &i)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Ingrese Tiempo que demora en entregar un paquete un camion")
	_, err2 := fmt.Scanf("%d\n", &i2)
	if err2 != nil {
		fmt.Println(err2)
	}
	for {
		time.Sleep((1) * time.Second)
		if CamionRetail1.Deruta == 0 {
			CamionRetail1.Deruta = 1
			go Send(&CamionRetail1, i, i2)

		}
		if CamionRetail2.Deruta == 0 {
			CamionRetail2.Deruta = 1
			go Send(&CamionRetail2, i, i2)

		}
		if CamionNormal.Deruta == 0 {
			CamionNormal.Deruta = 1
			go Send(&CamionNormal, i, i2)

		}

	}

}
