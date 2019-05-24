package main

import(
	"fmt"
	"sort"

)


func main(){


	quadrados := make(map[int]int, 15)
	numeros :=make([]int,0, len(quadrados))
	fmt.Println(numeros)

	for i:=1;i<=15;i++{
		quadrados[i] = i * i 
	}

	for n:= range quadrados{
		numeros = append(numeros, n)
	}

	sort.Ints(numeros)
	fmt.Println(numeros)


	for _ , numero := range numeros{
		quadrado := quadrados[numero]
		fmt.Printf("indice %d : %d\n", numero, quadrado)
	}
	

}