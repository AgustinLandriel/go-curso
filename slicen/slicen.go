package main

import "fmt"

func main() {

	numeros := []int{1, 2, 3}

	fmt.Println(numeros,len(numeros))

	//append agregar
	numeros = append(numeros,4)
	fmt.Println(numeros,len(numeros))
	numeros = append(numeros,5)
	fmt.Println(numeros,len(numeros))

	subNumero := numeros[:3]
	
	numeros [0] = 100
	
	fmt.Println(numeros,len(numeros))
	fmt.Println(subNumero,len(subNumero))


	//puntero
	//longitud
	//capacidad

	meses := [] string{"Enero","Febrero","Marzo"}

	fmt.Printf("len: %v,Cap: %v, %p \n",len(meses),cap(meses),meses)

	meses = append(meses,"Abril")
	fmt.Printf("len: %v,Cap: %v, %p \n",len(meses),cap(meses),meses)
}