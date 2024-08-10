package main

import (
	"log"
	"net/http"
	"Golang-Project-1/config"
	"Golang-Project-1/controllers/homecontroller"
	"Golang-Project-1/controllers/categorycontroller"
)



func main(){
	config.ConnectDB();

	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)










	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080",nil)
}

