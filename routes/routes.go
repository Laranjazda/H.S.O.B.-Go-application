package routes

import (
	"net/http"
	"web-service-application/controller"
)

func Init() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new-product", controller.NewProduct)
	http.HandleFunc("/insert", controller.SaveProduct)
	http.HandleFunc("/delete", controller.Delete)
}
