package golanglogger

import (
	"log"
	"time"
)

var loggerTime = time.Now()

func main() {

	return
}

func Kafka() time.Time {

	return loggerTime
}

func Cache() {
	return
}

// Loga as URI chamadas
func Logger() {
	log.Println("Logger", loggerTime)
	return
}

// Sample
// if rota.RequerAutenticacao {
// 	r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
// } else {
// 	r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
// }
func File() {
	log.Println("File", loggerTime)
	return
}
