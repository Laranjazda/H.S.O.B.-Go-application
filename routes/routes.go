package routes

import (
	"net/http"
	"web-service-application/controller"
)

func Init() {
	http.HandleFunc("/", controller.Index)
}
