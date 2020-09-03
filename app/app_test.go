package app

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
	"time"
)

func TestAppRun(t *testing.T) {
	Run()
}

func TestMongoInit(t *testing.T) {
	collection := Mongo.Database("kk").Collection("scrapy")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// res,err := collection.InsertOne(ctx,bson.M{"name":"run","date":"77"})
	res, err := collection.InsertOne(ctx, bson.M{"name": "run", "date": "66"})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(res.InsertedID)
}

func TestFindMap(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := crapMap.Find(ctx, bson.M{"key": "碧桂园"})
	if err != nil {
		log.Fatal(err)
	}
	// var result bson.A
	// cur.All(ctx,&result)
	// fmt.Println(result)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		fmt.Println(result["date"])
	}
}

func TestDelMapData(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	crapMap.DeleteMany(ctx, bson.M{"key": "万科"})
	crapMap.DeleteMany(ctx, bson.M{"key": "碧桂园"})
	crapData.DeleteMany(ctx, bson.M{"key": "万科"})
	crapData.DeleteMany(ctx, bson.M{"key": "碧桂园"})
}

// func TestMongoFind(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	cur, err := Collection.Find(ctx, bson.M{})
// 	if err != nil { log.Fatal(err) }
// 	defer cur.Close(ctx)
// 	for cur.Next(ctx) {
// 		var result bson.M
// 		err := cur.Decode(&result)
// 		if err != nil { log.Fatal(err) }
// 		// do something with result....
// 		fmt.Println(result)
// 	}
// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
