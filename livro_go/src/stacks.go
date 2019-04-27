package main


import(

	"fmt"
	"errors"
)


func main(){
	
	stack := Stack{}
	fmt.Println("Pilha criada com ", stack.Tamanho())
	fmt.Println("Vazia? ", stack.Vazia())
	
	stack.Empilhar("Go")
	stack.Empilhar(2009)
	fmt.Println("Tamanho da Pilha após empilhar 2 vezes: ", stack.Tamanho())

	stack.Desempilhar()
	fmt.Println("Tamanho da Pilha após desempilhar 1 vez: ", stack.Tamanho())
	

		
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


func (stack *Stack) Desempilhar() (interface{}, error) {
	
	if stack.Vazia(){
		return nil, errors.New("Pilha vazia")
	}

	
	valor:= stack.valores[stack.Tamanho()-1]
	stack.valores = stack.valores[:stack.Tamanho()-1]
	return valor, nil

}


