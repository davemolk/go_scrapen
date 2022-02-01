package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)


func main() {
	fName := "books.csv"
	f, e := os.Create(fName)
	if e != nil {
		log.Fatalf("Can't create file")
		return
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()

	// write CSV header
	writer.Write([]string{"title", "price", "available", "link"})

	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	c.OnHTML("ol.row li", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildAttr("img", "alt"),
			e.ChildText("p.price_color"),
			e.ChildText("p.instock"),
			e.ChildAttr("a", "href"),
		})
	})

	c.OnHTML("li.next", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Set error handler
	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})

	// pagination using num of pages
	// for i := 1; i <= 50; i++ {
	// 	url := fmt.Sprintf("https://books.toscrape.com/catalogue/page-%d.html", i)
	// 	c.Visit(url)
	// }


	c.Visit("https://books.toscrape.com/")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}

