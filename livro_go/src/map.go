package main

import(

	"fmt"

)

func main(){

	capitais := make( map[string]string)

	capitais["SC"] = "Florianopolis"
	capitais["RS"] = "Porto Alegre"
	capitais["PR"] = "Curitiba"

	fmt.Println(len(capitais))

	for estado, capital := range capitais{
		_ = estado
		fmt.Println(capital)
	}


}

