package main

import(

	"fmt"

)



func main(){

	type ListaGenerica []interface{}
	
	lista := ListaGenerica{"Cascavel"}
	lista = append(lista,"Londrina")
	lista = append(lista,"Irati")
	lista = append(lista,"Maringa")
	lista = append(lista,"Toledo")
	lista = append(lista,"Foz")
	lista = append(lista,"Curitiba")

	//remover := lista[1] 
	antes := lista[0:6]
	depois := lista[7:]
	nova_lista := append(antes, depois...)
	fmt.Println(antes)
	fmt.Println(depois)
	fmt.Println(nova_lista)
}