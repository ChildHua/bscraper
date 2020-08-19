package rcolly

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"math/rand"
)

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
		// link := e.Attr("href")
		link := e.ChildAttr("a", "href")
		// Print link
		m := map[string]interface{}{"name": e.Text, "link": link}
		data = append(data, m)
		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL.String())
	})

	for i := 0; i < 2; i++ {
		// Start scraping on https://hackerspaces.org
		vlink := fmt.Sprintf("http://www.baidu.com/s?rtt=1&bsst=1&cl=2&tn=news&word=%v&usm=3&rsv_idx=2&rsv_page=1&pn=%v", w, i*10)
		e := c.Visit(vlink)
		if e != nil {
			fmt.Println(e.Error())
		}
	}
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
