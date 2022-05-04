package main

import (
	"context"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func connectDB(collection string) mongo.Collection {
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://laranjazda:bros2011@h-s-o-b.5b97q.mongodb.net/HSOB?retryWrites=true&w=majority"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer client.Disconnect(ctx)

// 	// create database and collections
// 	hsobDB := client.Database("HSOB")

// 	return *hsobDB.Collection(collection)
// }

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantities  int32
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://laranjazda:bros2011@h-s-o-b.5b97q.mongodb.net/HSOB?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	// create database and collections
	hsobDB := client.Database("HSOB")
	productsDao := hsobDB.Collection("produtos")

	getTableProducts, err := productsDao.Find(ctx, bson.M{})
	if err != nil {
		panic(err.Error())
	}
	defer getTableProducts.Close(ctx)

	p := Product{}
	products := []Product{}

	for getTableProducts.Next(ctx) {
		var product bson.M

		if err = getTableProducts.Decode(&product); err != nil {
			log.Fatal(err)
		}
		p.Name = product["name"].(string)
		p.Description = product["description"].(string)
		p.Price = product["price"].(float64)
		p.Quantities = product["quantities"].(int32)

		products = append(products, p)
	}

	temp.ExecuteTemplate(w, "Index", products)
	// fmt.Println(products)
}
