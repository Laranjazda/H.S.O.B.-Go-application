package main

import (
	"net/http"
	"web-service-application/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.Init()
	http.ListenAndServe(":8000", nil)
}
