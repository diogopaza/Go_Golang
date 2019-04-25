package main

import(

	"fmt"
	"os"
	//"strings"

)

func main(){

	var palavras = os.Args[1:]
	

	//estatisticas := colherEstatisticas(palavras)
	//fmt.Println(estatisticas)
	//imprimir(estatisticas )
	colherEstatisticas(palavras)


}


func colherEstatisticas(palavras []string) /*map[string]int*/ {

	estatisticas := make(map[string]int)
	//fmt.Println(estatisticas)
	_ = estatisticas

	for _, palavra := range(palavras){
		var letraInicial = string(palavra[0])
		fmt.Println(string( letraInicial ))
	}


}


func imprimir(){


}