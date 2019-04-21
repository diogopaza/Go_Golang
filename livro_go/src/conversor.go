package main

import(

	"fmt"
	"os"	

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





	/*
	for i, v:= range valoresOrigem{


	}

	
	
	fmt.Println(unidadeOrigem)
	valorOrigem:= os.Args[len(os.Args)-2]
	fmt.Println(valorOrigem)
	*/




}