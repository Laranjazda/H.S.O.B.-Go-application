package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HsobDao struct {
}

func (connect *HsobDao) Collection(collection string) *mongo.Collection {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://laranjazda:bros2011@h-s-o-b.5b97q.mongodb.net/HSOB?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = dbClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return dbClient.Database("HSOB").Collection(collection)
}
