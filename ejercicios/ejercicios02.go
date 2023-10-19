package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var err error
var number int
var scan = bufio.NewScanner(os.Stdin)

func InputNumber() {
	iterar := func() {
		fmt.Println("Ingrese el primero número")
		if scan.Scan() {
			number, err = strconv.Atoi(scan.Text())
			fmt.Println("el numero que ingresaste es:", number)
		}
	}

	iterar()

	for err != nil {
		fmt.Println("El numero ingresado no esta permitido, debe ingresar otro número")
		if scan.Scan() {
			number, err = strconv.Atoi(scan.Text())
			fmt.Println("el numero que ingresaste es:", number)
		} else if err != nil {

		}

	}
}
