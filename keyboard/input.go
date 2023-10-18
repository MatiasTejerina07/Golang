package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var leyend string
var err error

func InputNumber() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese el primer numero")
	if scan.Scan() {
		number1, err = strconv.Atoi(scan.Text())
		fmt.Println(number1)
		if err != nil {
			panic("El dato ingresado no es un numero" + err.Error())
		}
	}

	fmt.Println("Ingrese el segundo  numero")
	if scan.Scan() {
		number2, err = strconv.Atoi(scan.Text())
		fmt.Println(number2)
		if err != nil {
			panic("El dato ingresado no es un numero" + err.Error())
		}
	}
}
