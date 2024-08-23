package main

import (
	"fmt"
)

var calculadora = map[string]func(int, int) int{
	"mas":      func(a int, b int) int { return a + b },
	"menos":    func(a int, b int) int { return a - b },
	"por":      func(a int, b int) int { return a * b },
	"dividido": func(a int, b int) int { return a / b },
}

func calcular(operacion string, a int, b int) {
	f, exists := calculadora[operacion]

	if !exists {
		fmt.Println("La operacion no es valida")
		return
	}

	resultado := f(a, b)
	fmt.Println("El resultado de la operacion de ", a, " ", operacion, " ", b, " es igual a: ", resultado)
}

var primer_valor int = 0
var funcionalidad string = ""
var segundo_valor int = 0

func main() {

	fmt.Print("Coloque el primer valor: ")
	fmt.Scanln(&primer_valor)

	fmt.Print("Que operacion desea realizar? (mas/menos/por/dividido)")
	fmt.Scanln(&funcionalidad)

	fmt.Print("Ingrese el segundo valor: ")
	fmt.Scanln(&segundo_valor)

	calcular(funcionalidad, primer_valor, segundo_valor)
}
