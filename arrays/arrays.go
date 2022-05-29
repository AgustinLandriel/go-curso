package main

import "fmt"

func main() {

	// 1er forma
	// var array [10] int

	//2da forma
	array2 := [10] string {"Agustin"}
	fmt.Println(array2)


	colores := [...] string {
		"rojo",
		"verde",
		"azul",
	} 

	fmt.Println(colores,len(colores))

	// indices definidos
	monedas := [...] string {
		0:"dolar",
		2:"Peso",
		3: "euros",
	}

	fmt.Println(monedas,len(monedas))

	//sub array
	subMoneda := monedas[:3]
	fmt.Println(subMoneda,len(subMoneda))
}