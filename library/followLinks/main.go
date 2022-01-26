package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Find and return all links
func main() {
	
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// callback on every a element with href
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q: %s", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://en.wikipedia.org/wiki/Go_(programming_language)")

}