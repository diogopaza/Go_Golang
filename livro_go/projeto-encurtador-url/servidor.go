package main

import(

	"fmt"
	"net/http"
	"log"
	"time"
)

var(

	porta int
	urlBase string
)

type Url struct{

	Id string
	Criacao time.time
	Destino string
}

func init(){
	porta = 8888
	urlBase = fmt.Sprintf("http://localhost:%d", porta)

}

func Encurtador(){

}

func Redirecionar(){

}

func main(){

		http.HandleFunc("/api", Encurtador)
		http.HandleFunc("/api", Redirecionar)

		log.Fatal(http.ListenAndServe(
			fmt.Sprintf(":%d", porta), nil))
		


}