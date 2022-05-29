package main

import "fmt"

func sumar () int {

	var a,b int 

	fmt.Println("Primer numero")
	fmt.Scanln(&a)
	fmt.Println("Segundo numero")
	fmt.Scanln(&b)

	resultado := a + b 
	
	return resultado
}

func calculoIva () float64 {

	var a float64 

	fmt.Println("Valor de producto")
	fmt.Scanln(&a)

	resultado := a + ( a * 21 ) / 100 
	
	return resultado
}

func main(){

	// r:=sumar()
	// fmt.Println(" La suma en total es : ",r)

	p := calculoIva()
	fmt.Println(" Total con IVA es : ",p)
}