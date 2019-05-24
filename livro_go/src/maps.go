package main

import(

	"fmt"
	"os"
	//"strings"

)

func main(){

	var palavras = os.Args[1:]
	

	estatisticas := colherEstatisticas(palavras)
	
	imprimir(estatisticas )
	


}


func colherEstatisticas(palavras []string) map[string]int {

	estatisticas := make(map[string]int)
	//fmt.Println(estatisticas)
	

	for _, palavra := range(palavras){
		var letraInicial = string(palavra[0])
		
		contador, encontrado := estatisticas[letraInicial]
		if encontrado{
			estatisticas[letraInicial] = contador + 1
		}else{
			estatisticas[letraInicial] = 1
		}
	}

	return estatisticas


}


func imprimir(estatisticas map[string]int){
	fmt.Println("Contagem de palavras iniciadas em cada letra:")
	for inicial, contador := range estatisticas{

		fmt.Printf("%s = %d\n", inicial, contador)
	}
		

}