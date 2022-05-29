package main

import "fmt"

func main() {

	var nombre string = "Agustin"

	// forma abreviada
	edad := 23

	// fmt.Println("Tu nombre es %s y tu edad es %d",nombre,edad)

	
	
	fmt.Printf("Tu nombre es %s y tu edad es %d\n",nombre,edad)

	
	mensaje := fmt.Sprintf("Tu nombre es %s y tu edad es %d",nombre,edad)

	fmt.Println(mensaje)

	apellido := ""
	fmt.Println("Ingrese apellido")
	fmt.Scanln(&apellido)
	fmt.Print("Apellido:",apellido)
}