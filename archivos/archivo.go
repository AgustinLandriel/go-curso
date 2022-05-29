package main

import (
	"fmt"
	"os"
)

func main() {

	//La funcion recover es parecido al panic pero no detiene el programa

	defer func () {
		if error := recover() ; error != nil {
			fmt.Print("Ups,al final el programa no finalizo de manera correcta")
		}
	} ()


	if file, error := os.Open("holasda.txt"); error != nil {
		panic("No fue posible obtener el archivo")
		//La funcion panic manda un error y detiene el programa
	} else {

		// Defer hace que una funcion se ejecute al final
		//Creo una funcion anonima que cierre archivo
		defer func() {
			fmt.Println("cerrando archivo")
			file.Close()
		} ()

		contenido := make ([] byte,254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)

	}

	
}