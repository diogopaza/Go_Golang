package main

import(

	"net/http"
	"fmt"
	"time"
	
)


func main(){

	
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
							fmt.Fprintf(w, "%s", time.Now().Format("2006-01-02 15:04:05"))	
							

					})
	
	fmt.Println("Rodando")
	http.ListenAndServe(":9000", nil)



}