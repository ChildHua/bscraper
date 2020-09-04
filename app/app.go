package app

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"test/rcolly"
	"test/utils"
	"time"
)

const DBNAME = "py"

var Mongo *mongo.Client
var crapMap, crapData *mongo.Collection
var KS []string
var sugarLogger *zap.SugaredLogger

func init() {
	KS = []string{"碧桂园", "万科", "保利", "金科"}
	MongoInit()
	InitLogger()
}

func Run() {
	go script()
	WebServer()
}

func script() {
	// 关键词设定
	// 定时
	sugarLogger.Info("begin script")
	now := time.Now()
	// 计算下一个零点
	next := now.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	t := time.NewTimer(next.Sub(now))
	// 抓取 && 存储
	for {
		for _, v := range KS {
			go scrap(v)
		}
		<-t.C
		sugarLogger.Info("begin crawl")
	}
}

func scrap(v string) {
	defer func() {
		if p := recover(); p != nil {
			sugarLogger.Panicf("crawl panic %v", p)
		}
	}()
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	sugarLogger.Info("begin crawl key:", v)

	d, _ := rcolly.Scrap(v)
	// res[v] = d
	djson, _ := json.Marshal(d)
	newM := bson.M{"key": v, "date": utils.TimeToDateStr(time.Now())}
	// _, err := crapMap.ReplaceOne(ctx, newM, newM)
	_, err := crapMap.DeleteOne(ctx, newM)
	_, err = crapMap.InsertOne(ctx, newM)
	if err != nil {
		sugarLogger.Panicf("mongo err: %v", err.Error())
	}
	_, err = crapData.DeleteOne(ctx, newM)
	_, err = crapData.InsertOne(ctx, bson.M{"key": v, "date": utils.TimeToDateStr(time.Now()), "data": string(djson)})
	if err != nil {
		sugarLogger.Panicf("mongo err: %v", err.Error())
	}
}
