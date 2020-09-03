package rcolly

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var crapMap, crapData *mongo.Collection

type DM map[string]interface{}

func Scrap(w string) ([]DM, error) {
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	// colly.AllowedDomains("baidu.com", "m.baidu.com"),
	)
	data := []DM{}
	// On every a element which has href attribute call callback
	c.OnHTML("h3", func(e *colly.HTMLElement) {

		link := e.ChildAttr("a", "href")

		m := map[string]interface{}{"name": e.Text, "link": link}
		data = append(data, m)

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL.String())
	})

	for i := 0; i < 1; i++ {
		// Start scraping on https://hackerspaces.org
		vlink := fmt.Sprintf("http://www.baidu.com/s?rtt=1&bsst=1&cl=2&tn=news&word=%v&usm=3&rsv_idx=2&rsv_page=1&pn=%v", w, i*10)
		e := c.Visit(vlink)
		if e != nil {
			fmt.Println(e.Error())
		}
	}
	return data, nil

}
