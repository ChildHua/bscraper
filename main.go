package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"test/rcolly"
	"time"
)

func main() {
	keywords := "保利,碧桂园,万科"
	r := gin.Default()
	r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("templates/*")
	jres := ""

	d := time.Duration(time.Minute * 10)

	t := time.NewTicker(d)
	// 每10分钟清理一次缓存

	go func() {
		for {
			<-t.C
			jres = ""
		}
	}()
	r.GET("/", func(c *gin.Context) {
		ks := strings.Split(keywords, ",")
		if len(jres) == 0 {
			res := make(rcolly.DM)
			for _, v := range ks {
				d, _ := rcolly.Scrap(v)
				res[v] = d
			}
			bres, _ := json.Marshal(res)
			jres = string(bres)
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"jsond": jres,
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
