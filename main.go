package main

import (
	"fmt"

	goroutines "github.com/MatiasTejerina07/Golang/goroutines"
)

func main() {
	/* texto, number := ejercicios.Ejercicio1("59")
	fmt.Println(texto, number)

	estado, text := variables.ConverText(20)
	fmt.Println(estado, text)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("Esto es linux")
	case "OS X":
		fmt.Println("Esto es mac")
	case "windows":
		fmt.Println("Esto es windows")
	default:
		fmt.Println("matias")
	} */

	/* iteraciones.Iterar() */
	/* iteraciones.JumpsFromTo() */
	/* iteraciones.AllExcept() */
	/* ejercicios.InputNumber() */
	//files.SaveTable()
	/* files.ReadText() */
	/* funciones.CallClosure() */
	/* funciones.Exponencia(2) */
	/* arreglos_slice.Capacidad() */
	/* mapas.ShowMaps() */
	/* Matias := new(modelos.Hombre)
	users.HumanosRespirando(Matias) */
	/* defer_panic.EjemploPanic() */
	go goroutines.MiNombreLentoooo("Matias Tejerina")

	fmt.Println("estoy aqui")
	var input string
	fmt.Scanln(&input)
	fmt.Println("EL VALOR DEL INPUTS ES :", input)
}
