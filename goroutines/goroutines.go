package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentoooo(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")

	for _, letras := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letras)
	}

	canal1 <- true

}
