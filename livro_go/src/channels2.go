package main

import(
	
	"fmt"
)


func produzir(c chan<- int){

	c <- 1
	c <- 2
	c <- 3
	close(c)
}

func main(){


	
}