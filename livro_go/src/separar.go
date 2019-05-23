package main

import(

	"fmt"
	
)



func separar(numeros []int, i,p chan<- int, pronto chan<- bool){

	for _, num := range numeros{
			
		if num % 2 == 0{
			p <- num
			
		}else{
			i <- num	
			
		}
		
		
	}

	pronto <- true
}


func main(){
	
	var pares []int
	var impares []int
	fim := false
	
	p,i:= make(chan int), make(chan int)
	pronto := make(chan bool)
	numeros := []int{3,9,15,12,29}
	go separar(numeros, i, p, pronto)

	for !fim{
		select{

		case n := <-i:
			impares = append(impares, n)
		case n:= <-p:
			pares = append(pares, n)
		case fim = <-pronto:
		}


	}

	fmt.Printf("Impares: %v | Pares:%v\n", impares, pares)


}