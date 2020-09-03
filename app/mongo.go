package app

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func MongoInit() {
	var err error
	Mongo, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Mongo.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	crapMap = Mongo.Database(DBNAME).Collection("crapMap")
	crapData = Mongo.Database(DBNAME).Collection("crapData")
}
