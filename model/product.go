package model

import (
	"context"
	"fmt"
	"log"
	"time"
	"web-service-application/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id          primitive.ObjectID
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

func DeleteProduct(name string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	hsobDao := mongodb.HsobDao{}
	productsDao := hsobDao.Collection("produtos")

	answer, err := productsDao.DeleteOne(ctx, bson.M{"name": &name})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer.DeletedCount)
}

func UpdateProduct(name string) Product {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	hsobDao := mongodb.HsobDao{}
	productsDao := hsobDao.Collection("produtos")

	answer, err := productsDao.Find(ctx, bson.M{"name": &name})
	if err != nil {
		log.Fatal(err)
	}

	product := Product{}

	for answer.Next(ctx) {
		var p bson.M

		if err = answer.Decode(&p); err != nil {
			log.Fatal(err)
		}
		product.Id = p["_id"].(primitive.ObjectID)
		product.Name = p["name"].(string)
		product.Description = p["description"].(string)
		product.Price = p["price"].(float64)
		product.Quantities = p["quantities"].(int32)
	}

	defer answer.Close(ctx)
	return product
}

func Update(name, description string, price float64, quantity int32) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	hsobDao := mongodb.HsobDao{}
	productsDao := hsobDao.Collection("produtos")

	answer, err := productsDao.UpdateOne(
		ctx,
		bson.M{"name": name},
		bson.D{
			{"$set", bson.D{
				{"name", name},
				{"description", description},
				{"price", price},
				{"quantities", quantity},
			}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer.ModifiedCount)

}
