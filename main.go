package main

import (
	"fmt"

	"github.com/MatiasTejerina07/Golang/variables"
)

func main() {
	estado, text := variables.ConverText(20)
	fmt.Println(estado, text)

}
