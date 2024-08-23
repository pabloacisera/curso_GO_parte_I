package main

import "fmt"

func main() {
	var fruta string

	fmt.Print("Escoga el producto que quiere consultar: ")
	fmt.Scanln(&fruta)

	switch fruta {
	case "coca-cola":
		fmt.Println("Precio de coca-cola 1 1/2 Lts: 1250.00")
		fmt.Println("Precio de coca-cola 1/2 Lts: 1250.00")
		fmt.Println("Precio coca-cola lata 450cc: 720.00")
	case "pepsi":
		fmt.Println("Precio Pepsi 1 1/4 Lts: 1125.00")
		fmt.Println("Precio de Pepsi 1/2 Lts: 1120.00")
		fmt.Println("Precio Pepsi lata 450cc: 610.00")
	default:
		fmt.Println("El produto no se encuentra disponible")
	}
}
