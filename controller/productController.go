package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"web-service-application/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := model.GetAllProduct()

	temp.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func SaveProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("price cannot be converted", err)
		}
		quantities, err := strconv.ParseInt(r.FormValue("quantities"), 10, 32)
		if err != nil {
			log.Println("quantities cannot be converted", err)
		}

		model.SaveNewProduct(name, description, price, int32(quantities))
	}

	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	model.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := model.EditProduct(id)
	temp.ExecuteTemplate(w, "Edit", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Price cannot be converted", err)
		}
		quantities, err := strconv.ParseInt(r.FormValue("quantities"), 10, 32)
		if err != nil {
			log.Println("Quantities cannot be converted", err)
		}

		model.Update(id, name, description, price, int32(quantities))
	}
	http.Redirect(w, r, "/", 301)
}
