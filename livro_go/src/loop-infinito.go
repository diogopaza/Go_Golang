package main

import(
	"fmt"
	"math/rand"
	"time"
)


func main(){

	rand.Seed(time.Now().UnixNano())
	n := 0

	for{
		n++
		i:= rand.Intn(4200)
		fmt.Println(i)
		if i%42 == 0{
			break
			
		}

	}

	fmt.Printf("Saida apoś %d interações\n", n)


	var i int
	externo:
	for{
		for i =0; i<10; i++{
			if i == 5{
				fmt.Println("indo pro break")
				break externo
			}
		}
	}

}