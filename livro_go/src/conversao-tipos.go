package main

import(

	"fmt"

)


type ListaDeCompras []string

func main(){

	lista := ListaDeCompras{"tomate","alface","feijao","carne"}
	imprimirLista(lista)
	slice := []string{ "tomate","alface","feijao","carne" } 
	imprimirSlice(slice)
    
}


func imprimirSlice(slice []string){

	fmt.Println("Slice: ", slice)

}

func imprimirLista(lista ListaDeCompras){

	fmt.Println("Lista de compras: ", lista)

}


