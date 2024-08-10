package categorycontroller

import (
	"time"
	"net/http"
	"Golang-Project-1/models/categorymodel"
	"Golang-Project-1/entities"
	"html/template"
	"strconv"
)


func Index(w http.ResponseWriter, r *http.Request){
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp,err := template.ParseFiles("views/category/index.html")
	if err != nil{
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		temp,err := template.ParseFiles("views/category/create.html")
		if err != nil{
			panic(err)
		}
		temp.Execute(w, nil)
	}


	if r.Method == "POST"{
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok:= categorymodel.Create(category); !ok{
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)

	}
}
func Edit(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		var category entities.Category
		id, err :=  strconv.ParseUint(r.FormValue("id"),10,64)
		category.Id = id 
		if err != nil{
			panic(err)
		}
		category.Name = r.FormValue("edit_name")
		category.UpdatedAt = time.Now()

		if ok:= categorymodel.Edit(category); !ok{
			temp, _ := template.ParseFiles("views/category/index.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)

	}

}
func Delete(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		var category entities.Category
		id, err := strconv.ParseUint(r.URL.Query().Get("id"),10,64)
		if err != nil {
			panic(err)
		}
		category.Id = id
		if ok:= categorymodel.Delete(category); !ok{
			temp, _ := template.ParseFiles("views/category/index.html")
			temp.Execute(w,nil)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	
	}
}