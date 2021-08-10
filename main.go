package golanglogger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
	_, fpath, _, ok := runtime.Caller(0)
	if !ok {
		err := errors.New("failed to get filename")
		panic(err)
	}
	filename := filepath.Base(fpath)

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pathFile := pwd
	ss := strings.Split(pathFile, "/")
	pathFolder := ss[len(ss)-1]

	fmt.Println("Folder =>", pathFolder)

	fmt.Println("File => ", filename)
}
