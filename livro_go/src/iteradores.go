package main

import(

	"fmt"
)

func main(){

	var a,b = 0,10
	
	for a < b{
		a = a + 1
		fmt.Println(a)
	}

	for i :=0; i<10; i++{
		a = a + 1
		fmt.Println(a)
	}
	
	numeros := []int{1,2,3,4,5}

	for indice, numero := range(numeros){
		_ = indice
		fmt.Println("O número da vez é: ", numero)
	}


}