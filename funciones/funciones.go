package funciones

import "fmt"

func Calculos() {
	suma := func(a int, b int) int {
		return a + b
	}
	result := suma(10, 25)
	fmt.Println(result)
}
