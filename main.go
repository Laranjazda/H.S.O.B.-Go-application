package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Colchão", "ortopédico", 1254.6, 5},
		{"Cama", "Casal", 2545.6, 2},
		{"Sofá", "tres lugares", 1654.6, 3},
		{"Camiseta", "Amarela", 25.65, 15},
	}
	temp.ExecuteTemplate(w, "Index", produtos)

}
