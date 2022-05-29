package main

import (
	"fmt"
	"strings"
)

func reverse (cadena string) string {

	//Split -> Convierte una cadena a array
	// Join -> Convierte un array a cadena

	arrayCadena := strings.Split(cadena,"")
	fmt.Println(arrayCadena)
	arrayInvertida:= make([] string,0)

	 for i := len(arrayCadena) -1; i >= 0; i-- {
		 arrayInvertida = append(arrayInvertida, arrayCadena[i])
	 }


	// fmt.Println(arrayInvertida)

	return strings.Join(arrayInvertida,"")
}

func esPalindromo(palabra string) bool {

	//oso
	
	palabra = strings.ToLower(palabra)
	palabra = strings.ReplaceAll(palabra," ","")
	// palabra = strings.ToUpper(palabra)
	fmt.Println(palabra)

	// fmt.Println(reverse(palabra))
	palabraInvertida := reverse(palabra)

	return palabra == palabraInvertida

}


func main ( ) {

	if esPalindromo("Ana"){
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}	
	
	
}