package main

import(

	"fmt"
	"time"
)

func executar(c chan<- bool){

	time.Sleep( 3 * time.Second)
	c <- true
}

func main(){

	fim := false

	c := make(chan bool, 1)
	go executar(c)
	fmt.Println("esperando")

	for !fim{
		select{
		case fim = <-c:
				fmt.Println("fim")
		case <- time.After(2 * time.Second):
			fmt.Println("Timeout")
			fim = true

		}


	}


}