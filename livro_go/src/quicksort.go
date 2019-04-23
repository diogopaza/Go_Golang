package main

import(

	"fmt"
	"os"
	"strconv"
	

)

func main(){

	var numeros_entrada = os.Args[1:]
	var numeros = make([]int, len(numeros_entrada))
	//fmt.Println(numeros)
	
	for i, n:= range numeros_entrada{

		numero, err := strconv.Atoi(n)
		if err != nil{
			fmt.Printf("%s não é um núemro válido \n", n)
			os.Exit(1)
		}

		numeros[i]=numero

	}

	fmt.Println(quicksort(numeros))
	
}


func quicksort(numeros []int)[]int{

	if len(numeros) <= 1{
		
		fmt.Println("array de uma posição ou zero")
		return numeros
	}

	n:= make([]int, len(numeros))
	copy(n, numeros)

	var indice_pivo = len(n)/2
	var pivo= n[indice_pivo]
	_ = pivo
	n = append(n[:indice_pivo], n[indice_pivo+1:]...)



	my_array:= make([]int, 5)
	my_array[0]= 24
	return n
}

