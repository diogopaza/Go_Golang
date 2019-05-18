package main

import(

	"fmt"
	"os"
)


func PrecoFinal( precoCusto float64) (
	precoDolar float64,
	precoReal float64, ){

	fatorLucro := 1.33
	taxaConversao := 3.9
	
	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaConversao

	return


}

func CriarArquivos(){



}

func main(){

	//fmt.Println( PrecoFinal(10) )
dirTemp := os.TempDir()
fmt.Println( dirTemp )


}