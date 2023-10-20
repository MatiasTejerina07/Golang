package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentoooo(nombre string) {
	letras := strings.Split(nombre, "")

	for _, letras := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letras)
	}

}
