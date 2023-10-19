package funciones

import "fmt"

func Table(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func CallClosure() {
	tabledel := 2
	MiTable := Table(tabledel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTable()) //Imprime la tabla del 2
	}
}
