package main

import(

	"fmt"
	"os"
	"strings"

)

func main(){

	var palavras = os.Args[1:]
	fmt.Println(palavras)


	estatisticas := colherEstatisticas(palavras)

	imprimir(estatisticas)



}


func colherEstatisticas(palavras []string) map[string]int {

	estatistias := make(map[string]int)


}


func imprimir(){


}