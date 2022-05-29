package main

import (
	"fmt"
	"strings"
)

// Closures
//Son funciones que retornan otras funcinones

//Esta funcion padre recibe a una funcion hija
func repeat (n int) func(cadena string) string  {

	return func (cadena string) string{
		return strings.Repeat(cadena,n)
		//Repite la cadena X veces
	}
}

func main () {

	repeat3 := repeat (3)
	fmt.Println(repeat3("Hola"))

	// 1er forma

	/* func ()  {
		fmt.Println("Hola")
	} ()
	*/

	// 2da forma en variable (mas usada)
	// mifun := func() {
	// 	fmt.Println("hOLA")
	// }

	// mifun()
}