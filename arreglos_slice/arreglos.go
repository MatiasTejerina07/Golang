package arreglos_slice

import "fmt"

var table [10]int = [10]int{10, 22, 55}

func MuestroArreglos() {
	table[7] = 33
	table[3] = 20

	tabla2 := [10]string{"hola", "mundo"}
	fmt.Println(tabla2)
	fmt.Println(table)

	for i := 0; i < len(table); i++ {
		fmt.Println(table)
	}
}
