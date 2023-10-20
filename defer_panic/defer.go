package defer_panic

import (
	"fmt"
	"log"
)

//el defer es una instruccion que nos da Go y nos permite configurar cual va hacer la ultima instruccion en ejercutarse cuando salga de la función
// se usa para manejar/ configurar las salidas (err, success, etc)
// se usa para cerrar recursos (archivos, etc)
// se usa para limpiar memoria (garbage collector)
// se usa para cerrar conexiones a bases de datos

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es un segundo mensaje")
	fmt.Println("Este es el último mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco)
		}
	}()

	a := 1

	if a == 1 {
		panic("esto es un error")
	}
}
