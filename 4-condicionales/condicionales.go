package main

import (
	"fmt"
	"strings"
)

func main() {
	var edad int
	var permiso string
	var acompañante string // Cambié el tipo de acompañante a string para manejar la entrada como "si/no"

	fmt.Print("Introduzca la edad del usuario: ")
	// Extraer el valor ingresado y guardarlo en la variable edad
	fmt.Scanln(&edad)

	fmt.Print("El usuario tiene permiso (si/no): ")
	fmt.Scanln(&permiso)

	fmt.Print("El usuario posee acompañante (si/no): ")
	fmt.Scanln(&acompañante)

	// Convertir la entrada de permiso y acompañante a minúsculas para manejar "si", "Si", "SI", etc.
	permiso = strings.ToLower(permiso)
	acompañante = strings.ToLower(acompañante)

	if edad >= 18 {
		fmt.Println("El sujeto es mayor de edad")
	} else {
		fmt.Println("El sujeto es menor de edad")
	}

	/*******************************************/

	// Evaluar las condiciones para permitir viajar
	if edad < 18 && permiso == "si" {
		fmt.Println("La persona puede viajar aunque sea menor porque cuenta con permiso")
	} else if edad < 18 && permiso == "no" && acompañante == "si" {
		fmt.Println("La persona puede viajar porque tiene acompañante aunque no cuenta con permiso")
	} else if edad < 18 && permiso == "no" && acompañante == "no" {
		fmt.Println("La persona no puede viajar porque no cuenta con permiso ni acompañante")
	} else if edad >= 18 {
		fmt.Println("La persona puede viajar por ser mayor de edad")
	} else {
		fmt.Println("La persona no puede viajar por ser menor de edad y no contar con permiso")
	}
}
