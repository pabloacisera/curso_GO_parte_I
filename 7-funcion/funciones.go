package main

import (
	"fmt"
)

var a int = 0
var b int = 0
var resultado int = 0

func suma(a int, b int) int {
	return a + b
}

func main() {
	fmt.Print("Indique el primer valor: ")
	fmt.Scanln(&a)
	fmt.Print("Indique el segundo valor: ")
	fmt.Scanln(&b)

	resultado = suma(a, b)
	fmt.Println("El resultado de la operacion es:", resultado)
}
