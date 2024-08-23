package main

import (
	"fmt"
	"strings"
)

func main() {

	var resultado int = 0
	var nums_pares = []int{}
	var nums_impares = []int{}

	for i := 0; i <= 19; i++ {
		resultado = i + 1
		fmt.Println(resultado)
	}

	for i := 0; i <= 19; i++ {
		if i%2 == 0 {
			nums_pares = append(nums_pares, i)
		} else {
			nums_impares = append(nums_impares, i)
		}
	}
	fmt.Println(nums_pares)
	fmt.Println(nums_impares)

	//tenemos tambien el foreach para trabajar con mapas

	var mi_mapa = map[string]string{
		"argentina": "buenos aires",
		"brasil":    "brasilia",
	}

	for k, v := range mi_mapa {
		fmt.Println("La capital de " + k + " es: " + v)
	}

	//para crear un while utilizamos un for infinito

	var pais string

	for {
		fmt.Print("Elige un país: ")
		fmt.Scanln(&pais)

		pais = strings.ToLower(pais)

		if pais == "peru" {
			fmt.Println("El usuario eligio Peru. La aplicacion finalizo")
			break
		} else {
			fmt.Println("El pais elegido es: " + pais + " y esta respuesta no es correcta")
		}
	}
}
