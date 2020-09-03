package rcolly

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"math/rand"
)

// type DM map[string]interface{}

func VScrap(w string) ([]DM, error) {
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	// colly.AllowedDomains("baidu.com", "m.baidu.com"),
	)
	data := []DM{}
	// On every a element which has href attribute call callback
	c.OnHTML("p[class]", func(e *colly.HTMLElement) {
		class := e.Attr("class")
		if class == "tip" {
			fmt.Println(e.Text)
		}
		link := e.ChildAttr("a", "href")
		// Print link
		m := map[string]interface{}{"name": e.Text, "link": link}
		data = append(data, m)
	})

	c.OnHTML("div[class]", func(e *colly.HTMLElement) {
		h := "https://cn.vuejs.org/"
		class := e.Attr("class")
		if class == "guide-links" {
			hrefs := e.ChildAttrs("a", "href")
			if len(hrefs) == 1 {
				h += hrefs[0]
			} else {
				h += hrefs[1]
			}
			c.Visit(h)
			// fmt.Println(hrefs)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL.String())
	})

	vlink := "https://cn.vuejs.org/v2/guide/installation.html"
	c.Visit(vlink)
	// for i := 0; i < 2; i++ {
	// 	// Start scraping on https://hackerspaces.org
	// 	vlink := fmt.Sprintf("http://www.baidu.com/s?rtt=1&bsst=1&cl=2&tn=news&word=%v&usm=3&rsv_idx=2&rsv_page=1&pn=%v", w, i*10)
	// 	e := c.Visit(vlink)
	// 	if e != nil {
	// 		fmt.Println(e.Error())
	// 	}
	// }
	return data, nil
	// bdata,_ := json.Marshal(data)
	// return string(bdata),nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
