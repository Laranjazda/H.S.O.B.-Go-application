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
	Id          primitive.ObjectID `bson:"_id" json:"id, omitempty"`
	Idstr       string
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantities  int32   `json:"quantities"`
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
		/*CONVERT OBJECID TO STRING*/
		p.Id = product["_id"].(primitive.ObjectID)
		p.Idstr = p.Id.Hex()
		p.Name = product["name"].(string)
		p.Description = product["description"].(string)
		p.Price = product["price"].(float64)
		p.Quantities = product["quantities"].(int32)

		products = append(products, p)
	}
	defer getProducts.Close(ctx)
	return products
}

func DeleteProduct(id string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	hsobDao := mongodb.HsobDao{}
	productsDao := hsobDao.Collection("produtos")

	/*CONVERT STRING TO OBJECTID*/
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Panic(err.Error())
	}

	answer, err := productsDao.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer.DeletedCount)
}

func EditProduct(id string) Product {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	hsobDao := mongodb.HsobDao{}
	productsDao := hsobDao.Collection("produtos")

	/*CONVERT STRING TO OBJECTID*/
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Panic(err.Error())
	}

	answer, err := productsDao.Find(ctx, bson.M{"_id": objId})
	if err != nil {
		log.Fatal(err)
	}

	product := Product{}

	for answer.Next(ctx) {
		var p bson.M

		if err = answer.Decode(&p); err != nil {
			log.Fatal(err)
		}
		/*CONVERT OBJECID TO STRING*/
		product.Id = p["_id"].(primitive.ObjectID)
		product.Idstr = product.Id.Hex()
		product.Name = p["name"].(string)
		product.Description = p["description"].(string)
		product.Price = p["price"].(float64)
		product.Quantities = p["quantities"].(int32)
	}

	defer answer.Close(ctx)
	return product
}

func Update(id, name, description string, price float64, quantity int32) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	hsobDao := mongodb.HsobDao{}
	productsDao := hsobDao.Collection("produtos")

	/*CONVERT STRING TO OBJECTID*/
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Panic(err.Error())
	}

	answer, err := productsDao.UpdateOne(
		ctx,
		bson.M{"_id": objId},
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
