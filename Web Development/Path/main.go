package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Sandal</h1>")

}


func main(){
	http.HandleFunc("/", helloWorldPage)
	http.ListenAndServe("",nil)
}