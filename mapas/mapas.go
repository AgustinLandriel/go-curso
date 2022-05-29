package main

import "fmt"

func main() {

	//los mapas son como diccionarios de python
	dias := make(map[int] string)
	
	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "martes"
	dias[3] = "miercoles"
	dias[4] = "jueves"
	dias[5] = "vienres"
	dias[6] = "sabado"
	// fmt.Println(dias[3])

	// delete(dias,2)
	// fmt.Println(dias)

	//nuevo mapas
	estudiantes := make(map[string][]int) 

	estudiantes ["agustin"] = [] int {13,15,16}
	estudiantes ["matias"] = [] int {1,3,66}
	// clave | valor
	fmt.Println(estudiantes)

	fmt.Println(estudiantes["agustin"][1])
}