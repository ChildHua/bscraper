package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
)

func WebServer() {
	r := gin.Default()
	r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("templates/*")

	r.GET("/map/", Map)
	r.GET("/detail/:kd", Detail)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func Map(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := make(map[string][]string)
	for _, v := range KS {

		mapData, err := crapMap.Find(ctx, bson.M{"key": v})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		for mapData.Next(ctx) {
			var result bson.M
			err := mapData.Decode(&result)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			if result["date"] != nil {
				s := result["date"].(string)
				res[v] = append(res[v], s)
			}
		}
	}
	rb, _ := json.Marshal(res)
	fmt.Println("res", string(rb))
	// c.JSON(http.StatusOK,res)
	c.HTML(http.StatusOK, "map.tmpl", gin.H{
		"title": "Main website",
		"key":   KS,
		"data":  res,
	})
}

func Detail(c *gin.Context) {
	kd := strings.Split(c.Param("kd"), ",")
	k, d := kd[0], kd[1]
	datac, err := crapData.Find(c, bson.M{"key": k, "date": d})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var result bson.M
	datac.Next(c)
	err = datac.Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "detail.tmpl", gin.H{
		"title": "Main website",
		"key":   k,
		"date":  d,
		"data":  result["data"],
	})
}
