package controller

import (
	"net/http"
	"text/template"
	"web-service-application/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := model.GetAllProduct()

	temp.ExecuteTemplate(w, "Index", products)
	// fmt.Println(products)
}
