package main

import(
	"fmt"
)

func main(){

	/*
	fib :=[]int{1,1,2,3,5,8,13}
	fmt.Println(fib)
	fmt.Println(fib[:3])
	fmt.Println(fib[2:])
	fmt.Println(fib[:])
	

	original := []int{1,2,3,4,5}
	fmt.Println("Original: ", original)

	novo := original[1:3]
	fmt.Println("Novo: ", novo)

	original[2] = 13
	fmt.Println("Original pós modificação: ", original)
	fmt.Println("Novo após modificação: ", novo)
	

	a := []string{"Paulo","almoça","em","casa","diariamente"}
	b := a[:len(a)-1]
	c := b[:len(b)-1]
	d := c[:len(c)-1]
	e := d[:len(d)-1]
	
	e[0]="Tiago"
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n", a,b,c,d,e)
	
	s:= make([]int, 0)
	s= append(s,23)
	fmt.Println(s)
	
	s:=[]int{23,24,25}
	n := []int{22}
	n= append(n,s...)
	fmt.Println(n)
	*/

	/*
	s:= []int{11,12,13,16,17,18}
	v:= []int{14,15}
	final := s[3:6]
	s = append(s[:3],v...)
	s = append(s,final...)
	fmt.Println(s)
	

	s:= []int{11,12,13,16,17,18}
	v:= []int{14,15}
	_ = v
	s = append(s[:3], append(v, s[3:]...)...)
	fmt.Println(s)
	

	s:=[]int{10,20,30,40,50,60}
	//final := make([]int,0)
	//inicio := make([]int,0)
	//s = s[1:]
	//s = s[:3]
	//inicio = s[:2]
	//final = s[4:]//pega as duas primeiras posições do array
	s = append(s[:2],s[4:]...)
	fmt.Println(s)//10 20 50 60 remover 30 40 
	*/

	numeros := []int{1,2,3,4,5}
	dobros:= make([]int, len(numeros))

	copy(dobros, numeros)

	for i:= range(dobros){
		dobros[i] *= 2

	}

	fmt.Println(numeros)
	fmt.Println(dobros)

}