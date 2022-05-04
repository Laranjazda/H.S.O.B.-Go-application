package model

import (
	"context"
	"log"
	"time"
	"web-service-application/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantities  int32
}

func GetAllProduct() []Product {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	hsobDao := mongodb.HsobDao{}
	productsDao := hsobDao.Collection("produtos")

	getTableProducts, err := productsDao.Find(ctx, bson.M{})
	if err != nil {
		panic(err.Error())
	}

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
	defer getTableProducts.Close(ctx)
	return products
}
