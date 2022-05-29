package main

import "fmt"

	
func saludar ( nombre string,apellido string) {

	fmt.Println("Hola, ",nombre,apellido)
}

func sumar ( a,b int) int {

	return a + b
}
func main(){

	saludar("agustin","landriel")
	resultado := sumar(10,20)
	fmt.Println("La suma es " , resultado)
}