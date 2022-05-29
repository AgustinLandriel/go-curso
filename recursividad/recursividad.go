package main

import "fmt"

func factorial(n int) int {

	//Recursividad, se multiplica un numero por sus
	//antecesores,ej: 3 seria 3*2*1
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {

	// 3 ->  1 2 3
	fmt.Println(factorial(5))

}