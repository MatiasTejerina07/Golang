package ejercicios

import (
	"strconv"
)

func Ejercicio1(texto string) (int, string) {
	convertNumber, _ := strconv.Atoi(texto)
	if convertNumber > 100 {
		return convertNumber, "El número es mayor a 100"
	} else {
		return convertNumber, "El número es menor a 100"
	}
}
