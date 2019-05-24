package main

import(
	
	"fmt"

)


func main(){

	//array de inteiroscom tamamho 2
	var meu_array [2]int
	fmt.Println(meu_array)

	//arrays cujos valores são arrays
	var multiA [2][2]int
	fmt.Println(multiA)

	multiA [0][0], multiA[0][1] = 3, 5 
	fmt.Println(multiA)


	//slices
	var a []int
	primos := []int{2,3,5,7,11,13}
	nomes := []string{}

	fmt.Println(a)
	fmt.Println(primos)
	fmt.Println(nomes)

	b:= make([]int, 10)
	fmt.Println(b, len(b), cap(b))

	c:= make([]int,10,20)
	fmt.Println(c, len(c), cap(c))

	




}



