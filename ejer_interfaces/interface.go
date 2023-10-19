package ejer_interfaces

import (
	"fmt"

	"github.com/MatiasTejerina07/Golang/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s y estoy respirando \n", hu.Sexo())
}
