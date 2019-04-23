package main

import(

	"fmt"
	"os"
	//"strconv"

)

func main(){

	var numeros_entrada = os.Args[1:]
	

	for i, v := range(numeros_entrada){

		fmt.Printf("Numero %s no indice %d \n",v,i)

	} 

}

