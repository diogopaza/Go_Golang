package main

import "fmt"



	type person struct{
		name string
		age int

	}



func main(){

	fmt.Println(person{"Diogo", 35})
	s := person{name: "Sean", age: 50}


	fmt.Println(s.name)
	sp := &s

	fmt.Println(sp.age)


}





