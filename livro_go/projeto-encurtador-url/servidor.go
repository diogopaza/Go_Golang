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
	Criacao time.Time
	Destino string
}

type Headers map[string]string

func init(){
	porta = 8888
	urlBase = fmt.Sprintf("http://localhost:%d", porta)
	fmt.Println(urlBase)

}

func encurtador(w http.ResponseWriter, r *http.Request){

	if r.Method != "GET"{
		responderCom(w, http.StatusMethodNotAllowed, Headers{
			"Allow":"Post",			
		})
	}	

	result:= extrairUrl(r)
	fmt.Println(result)
		
		
		
		
	

}

func redirecionar(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Estou no redirecionar"))

}

func responderCom(w http.ResponseWriter,
					status int,
					headers Headers,){
						
						for k, v:= range headers{
							w.Header().Set(k,v)
						}
						
						w.WriteHeader(status)
}


func extrairUrl(r *http.Request)string{

	fmt.Println("extrair")
	fmt.Println(r.ContentLength)
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
	

}





func main(){

	
		http.HandleFunc("/encurtar", encurtador)
		http.HandleFunc("/r", redirecionar)
		
		log.Fatal(http.ListenAndServe(
			fmt.Sprintf(":%d", porta), nil))
		


}