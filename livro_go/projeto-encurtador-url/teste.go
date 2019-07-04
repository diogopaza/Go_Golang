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
	fmt.Printf("ContentLength:%v\n",contentlength)
	url := make([]byte, r.ContentLength, r.ContentLength )
	r.Body.Read(url)
	fmt.Println("Body", url)

} 

func main(){

	
	http.HandleFunc("/", length)

	http.ListenAndServe(":8000", nil)

}