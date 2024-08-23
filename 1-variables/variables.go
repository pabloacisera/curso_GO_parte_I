package main

import (
	"fmt"
)

func main() {

	miVariable := "pablo cisera"

	fmt.Println("Este es el valor de la variable:", miVariable)

	//esto es un comentario

	//las variables tambien pueden utilizar var + nombre + tipo de dato

	var variable_numerica int = 37

	fmt.Println(variable_numerica)

	//se puede realizar operaciones sobre estas variables

	fmt.Println("mi nombre es: "+miVariable, "y mi edad es: ", +variable_numerica)

	//la declaracion corta es igual a decir var

	asig_corta := "soy programador"

	fmt.Println(asig_corta)

	// los numero puede ser inferidos con la asignacion corta o pueden ser declarados con tipo y magnitud

	var año int16 = 2024
	var mes int8 = 9
	var dia int8 = 22

	fmt.Println("fecha: ", +año, "/", +mes, "/", +dia)

	//utilizamos float para decimales

	var numero float32 = 15
	var numero_dos float32 = 6542
	resultado := numero + numero_dos
	fmt.Println(resultado)

}
