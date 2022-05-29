package main

import "fmt"

func main() {

	/*
	* Descuento de 10% hasta 100 de consumo
	* Descuento de 20% hasta 200 de consumo
	* Descuento de 30% mayor 200 de consumo
	* IVA 21 %
	*/


	var consumo, descuento float64
	var datosDescuento string

	// entrada de datos
	fmt.Println("ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0  {
		if consumo <=100 {
			datosDescuento = "10%"
			descuento = consumo * 0.10 
		} else if consumo > 100 && consumo <=200 {
			datosDescuento = "20%"
			descuento = consumo * 0.20
		} else if consumo > 200 {
			datosDescuento = "30%"
			descuento = consumo * 0.30 
		} else {
		fmt.Println("Error al ingresar consumo")}
	}


	//OPERACIONES
	var montoDesc = consumo - descuento
	var iva = montoDesc * 0.21
	var total = montoDesc + iva

	fmt.Println("-----Factura de Consumo-----")
	fmt.Println("Descuento que se aplica: ",datosDescuento)
	fmt.Println("Monto descontado: ",montoDesc)
	fmt.Printf("IVA: %2.f\n",iva)
	fmt.Println("Total a pagar: ",total)
}