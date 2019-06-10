package main

import(

	"fmt"
	"net/http"

)

func length(w http.ResponseWriter, r *http.Request){

	res, err := http.Head("https://www.bol.com.br/")
	if err != nil {
    	panic(err)
	}
	contentlength:=res.ContentLength
	fmt.Printf("ContentLength:%v",contentlength)

} 

func main(){

	
	http.HandleFunc("/", length)

	http.ListenAndServe(":8000", nil)

}