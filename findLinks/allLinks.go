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

	// find and print all links
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		fmt.Println(links)
	})

	c.Visit("https://en.wikipedia.org/wiki/Go_(programming_language)")

}