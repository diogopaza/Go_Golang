package main


import(

	"fmt"
)


func main(){

	stack := Stack{}
	fmt.Println("Pilha criada com ", stack.Tamanho())
	fmt.Println("Vazia? ", stack.Vazia())
	
	stack.Empilhar("Go")
	stack.Empilhar(2009)
	//stack.Empilhar(3.14)
	//stack.Empilhar("Fim")

	fmt.Println("Tamanho da Pilha após empilhar 2 vezes: ", stack.Tamanho())
	//fmt.Println("Vazia? ", stack.Vazia())

	numeros := make([]int, 0)
	numeros = append(numeros, 5)
	

	/*
	var meuArray [10]int
	fmt.Println(meuArray)
	*/

	
}

type Stack struct{
	
	valores []interface{}
	
}

func (stack Stack) Tamanho() int{
	return len(stack.valores)
}

func (stack Stack) Vazia() bool{

	return stack.Tamanho() == 0

}

func (stack *Stack) Empilhar( valor interface{}){
	stack.valores = append(stack.valores, valor)

}
