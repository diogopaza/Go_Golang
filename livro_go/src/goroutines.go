package main

import(

	"fmt"
	"time"
)


func imprimir(n int){

	for i:=0;i<3;i++{
		fmt.Println("%d", n)
		time.Sleep(800 * time.Millisecond)

	}

}

func dormir(){
	fmt.Println("Funcao dormindo 5 segundos")
	time.Sleep( 5 * time.Second)
	fmt.Println("Finalizando função dormir")
}


func main(){

	go dormir()
	fmt.Println("Main dormindo por 3 segundos")
	time.Sleep( 3 * time.Second)
	fmt.Println("main finalizada")


}