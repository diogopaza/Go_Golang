package main

import(

	"fmt"
)


type Arquivo struct{

	nome string
	tamanho float64
	caracteres int
	palavras int

}

func ( arq *Arquivo) TamanhoMedioDePalavra() float64{
	return float64(arq.caracteres) / float64(arq.palavras)
}

func main(){

	
	arquivo := Arquivo{"artigo.txt", 12.68, 12986, 220}
	//arquivoNovo := Arquivo{ nome:"novo.txt", palavras: 2525}
	//fmt.Println( arquivo )
	//fmt.Println( arquivoNovo )
	


	ponteiroArquivo := &Arquivo{tamanho:7.29, nome:"arqivoPonteiro"}
	fmt.Printf("%s %.2fKB\n",
				ponteiroArquivo.nome,
				ponteiroArquivo.tamanho,	
	)

	ponteiroArquivo.tamanho = 50
	fmt.Printf("%s %.2fKB\n",
				ponteiroArquivo.nome,
				ponteiroArquivo.tamanho,	
	)



	fmt.Printf("Tamanho médio das palavras: %.2f\n", arquivo.TamanhoMedioDePalavra() )



}