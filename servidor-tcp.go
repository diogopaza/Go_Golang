package main

import(

	"fmt"
	"net"
	"os"
	"bufio"
	"strings"

)


func main(){

	fmt.Println("Servidor aguardando conexões...")

	ln, err := net.Listen("tcp", ":8081")
	if err != nil{
		fmt.Println("Erro ao conectar")
		os.Exit(3)
	}

	conexao, err := ln.Accept()
	if err != nil{
		fmt.Println("Erro ao receber conexao")
		os.Exit(3)
	}

	defer ln.Close()

	fmt.Println("Conexao aceita")
	for{

		mensagem, err := bufio.NewReader(conexao).ReadString('\n')
		if err != nil{
			fmt.Println(err)
			os.Exit(3)
		}
	

		fmt.Println("Mensagem recebeida: ", string(mensagem) )

		novaMensagem := strings.ToUpper(mensagem)
		conexao.Write([]byte(novaMensagem +"\n") )
	}


}