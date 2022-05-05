package model

import (
	"context"
	"fmt"
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

func SaveNewProduct(name, description string, price float64, quantities int32) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	hsobDao := mongodb.HsobDao{}
	productsDao := hsobDao.Collection("produtos")
	answer, err := productsDao.InsertOne(ctx,
		bson.D{
			{Key: "name", Value: &name},
			{Key: "description", Value: &description},
			{Key: "price", Value: &price},
			{Key: "quantities", Value: &quantities},
		},
	)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(answer.InsertedID)

}
func GetAllProduct() []Product {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	hsobDao := mongodb.HsobDao{}
	productsDao := hsobDao.Collection("produtos")

	getProducts, err := productsDao.Find(ctx, bson.M{})
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for getProducts.Next(ctx) {
		var product bson.M

		if err = getProducts.Decode(&product); err != nil {
			log.Fatal(err)
		}
		p.Name = product["name"].(string)
		p.Description = product["description"].(string)
		p.Price = product["price"].(float64)
		p.Quantities = product["quantities"].(int32)

		products = append(products, p)
	}
	defer getProducts.Close(ctx)
	return products
}
