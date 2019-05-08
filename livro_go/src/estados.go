package main

import(

	"fmt"
)


type Estado struct{

	nome string
	populacao int
	capital string

}

func main(){
	/*
	estados := make( map[string]Estado )
	estados["GO"] = Estado{nome:"Goias",populacao:99999, capital:"Goiania"}
	estados["PR"] = Estado{nome:"Parana",populacao:99999, capital:"Curitiba"}
	estados["SC"] = Estado{nome:"Santa Catarina",populacao:99999, capital:"Florianopolis"}
	estados["SP"] = Estado{nome:"Sao Paulo",populacao:99999, capital:"Sao Paulo"}
	estados["RJ"] = Estado{nome:"Rio de Janeiro",populacao:99999, capital:"Rio de Janeiro"}
	estados["MT"] = Estado{nome:"Mato Grosso",populacao:99999, capital:"Cuiaba"}

	//fmt.Println(estados)

	parana := estados["PR"]
	fmt.Println(parana)
	fmt.Println(estados["RN"])

	saoPaulo, encontrado := estados["SP"]
	if encontrado {
		fmt.Println(saoPaulo)
	}
	*/

	idades := make(map[string]int)
	idades["Diogo"] = 34
	idades["Beto"] = 59
	fmt.Println(idades)
	delete(idades,"Beto") 
	fmt.Println(idades)


}