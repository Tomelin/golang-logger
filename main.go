package golanglogger

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var loggerTime = time.Now()

type LoggerErr struct {
	APIName    string
	file       string
	err        error
	statusCode *string
	time       time.Time
}

func main() {

	return
}

func Kafka() time.Time {

	return loggerTime
}

func Cache() {
	return
}

// Loga os erros da API.  Parametros necessários APINmae, nome do arquivo, erro, status code
func Logger(APIName string, file string, err error, statusCode string) {
	log.Println("Logger....", APIName, file, err, statusCode, time.Now())
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

	return
}

func FileName(fName string) {
	fpath := fName
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

	return
}

func LoggerFile(fName string) {
	filename := filepath.Base(&fName)

	pwd, err := filepath.Dir(&fName)
	if err != nil {
		os.Exit(1)
	}

	ss := strings.Split(pwd, "/")
	pathFolder := ss[len(ss)-1]

	fmt.Println("Folder =>", pathFolder)

	fmt.Println("File => ", filename)
	err := filepath.Walk(filepath.Dir(&fName), getModuleName)
	if err != nil {
		log.Fatal(err)
	}
	return

}

func getModuleName(path string, info os.FileInfo, err error) string {

	if err != nil {
		return nil
	}

	pathSplit := strings.Split(path, "/")
	fName := pathSplit[len(pathSplit)-1]
	if fName == "go.mod" {
		oFile, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer oFile.Close()

		scanner := bufio.NewScanner(oFile)
		r, err := regexp.Compile("module")

		if err != nil {
			log.Fatal(err)
		}

		for scanner.Scan() {
			if r.MatchString(scanner.Text()) {
				oSplit := strings.Split(scanner.Text(), " ")
				return oSplit[1]
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
