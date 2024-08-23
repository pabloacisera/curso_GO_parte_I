package main

import "fmt"

func main() {

	fmt.Println("Capitulo Maps/diccionarios")

	var miMapa = map[string]string{
		"Colombia":  "Bogota",
		"Argentina": "Buenos Aires",
		"Brasil":    "Brasilia",
	}

	fmt.Println(miMapa) //los ordena automaticamente

	fmt.Println("Capital de Argentina: ", miMapa["Argentina"])

	miMapa["Peru"] = "Quito"

	fmt.Println(miMapa)

	/***********************************************************/
	//quitar elementos del map

	delete(miMapa, "Peru")
	fmt.Println(miMapa)
}
