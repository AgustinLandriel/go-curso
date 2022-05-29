package main

import "fmt"

// Struct persona

//Clase persona
type Persona struct {

	//Atributos
	nombre string
	edad int
}

//metodo
func (p *Persona) datos()  {
	
	fmt.Printf("edad: %d\nNombre: %s\n",p.edad,p.nombre)
}

func (p *Persona) setNombre (nuevoNombre string) {

	p.nombre = nuevoNombre
}

//Herencia
type Empleado struct {

	Persona
	sueldo float64

}


func main () {

	//Construyo el Objeto
	agustin := Persona {"Agustin",22}
	agustin.setNombre("chucho")

	agustin.datos()

	// Herencia

	em1 := Empleado {sueldo: 30.6}

	em1.nombre = "Pedro"
	em1.edad = 30
	em1.datos()
	fmt.Println(em1.sueldo)
}