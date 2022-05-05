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
		var k string = r.FormValue("price")
		var h string = r.FormValue("quantities")
		price, err := strconv.ParseFloat(k, 64)
		if err != nil {
			log.Println("price cannot be converted", err)
		}
		quantities, err := strconv.ParseInt(h, 10, 32)
		if err != nil {
			log.Println("quantities cannot be converted", err)
		}

		model.SaveNewProduct(name, description, price, int32(quantities))
	}

	http.Redirect(w, r, "/", 301)

}
