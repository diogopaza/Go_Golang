package main

import(

	"fmt"
	"os"	
	"strconv"

)

func main(){

	var unidadeDestino string
	
	var valoresOrigem = os.Args[1 : len(os.Args)-1]
	
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	fmt.Println(valoresOrigem)
	unidadeOrigem:= os.Args[len(os.Args)-1]

	if unidadeOrigem == "celsius"{
		unidadeDestino = "fahrenheit"
	}else if unidadeOrigem == "quilometros"{
		unidadeDestino = "milhas"
	}else{
		fmt.Println("%s não é uma unidade conhecida", unidadeDestino)
		os.Exit(1)
	}

	for  i, v:= range valoresOrigem{

		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil{
			fmt.Println("O valor %s na posição %d não é um número valido!\n", v, i)
			os.Exit(1)
		}

		
		var valorDestino float64 
		_ = valorDestino
		if unidadeOrigem == "celsius"{
			
			valorDestino = valorOrigem * 1.8 + 32
			fmt.Printf("%.2f %s = %.2f %s\n ", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
			
		}else{
			
			valorDestino = valorOrigem / 1.60934
			fmt.Printf("%.2f %s = %.2f %s\n ", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
			
		}

		
		
	}




	
	/*
	fmt.Println(unidadeOrigem)
	valorOrigem:= os.Args[len(os.Args)-2]
	fmt.Println(valorOrigem)
	*/




}