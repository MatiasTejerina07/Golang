package mapas

import "fmt"

func ShowMaps() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["MEXICO"] = "DF"
	paises["ARGENTINA"] = "BS AS"
	fmt.Println(paises)
	fmt.Println(paises["ARGENTINA"])

	campeonato := map[string]int{"BARCA": 39, "RIVER": 41, "BOCA": 41, "REAL MADRID": 66}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("EQUIPO %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	delete(campeonato, "BOCA")

	fmt.Println(campeonato)

	puntaje, existe := campeonato["BARCA"]

	fmt.Printf("El puntaje que tiene es %d y el equipo existe %t", puntaje, existe)

}
