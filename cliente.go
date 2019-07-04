package main

import(

	"fmt"
	"net"
	"bufio"
	"os"
)


func main(){

	conexao, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil{
		fmt.Println("Erro ao conectar")
	}

	for{
		leitor := bufio.NewReader(os.Stdin)
		fmt.Println("texto a ser enviado: ")
		texto, err := leitor.ReadString('\n')
		if err != nil{
			fmt.Println("Erro ao enviar mensagem")
		}

	

		fmt.Fprintf(conexao, texto +"\n")

		respostaServidor, err := bufio.NewReader(conexao).ReadString('\n')
		if err != nil{
			fmt.Println("Erro ao receber resposta do servidor")
		}

		fmt.Println("Resposta do servidor: ", respostaServidor)
	}	

}
	