package main

import "fmt"

func main() {
	//arreglos y slice

	var lista = [2]string{}
	lista[0] = "argentina"
	lista[1] = "colombia"

	var lista2 = [2]int8{8, 12}

	fmt.Println(lista)
	fmt.Println(lista2)

	/*********los datos del array pueden modificarse*********/

	//slice es igual a un arreglo pero sin tamaño definido:

	data := []string{"pablo", "andres"}
	fmt.Println(data)
	data = append(data, "cisera") //agregamos el elemento al final de la lista
	fmt.Println("El nombre completo del usuario es: ", data)

	//podemos obtener rangos, desde la posicion hasta la posicion

	nombres := data[0:2]
	fmt.Println("El nombre del usuarios es: ", nombres)

	apellido := data[2:]
	fmt.Println("El apellido del usuario es: ", apellido)

	//
}
