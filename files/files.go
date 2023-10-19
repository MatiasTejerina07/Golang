package files

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/MatiasTejerina07/Golang/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func SaveTable() {
	var text string = ejercicios.InputNumber()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, text)
	archivo.Close()
}

func AddTable() {
	var text string = ejercicios.InputNumber()
	if !Append(filename, text) {
		fmt.Println("Error al agregar el archivo")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error al WriteString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func ReadText() {
	datos, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}
	fmt.Println(string(datos))
}
