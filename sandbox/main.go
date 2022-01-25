package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("p.introduction a strong", func(e *colly.HTMLElement) {
		// just p.introduction
		// fmt.Println(e.Text)
		// link := e.ChildAttrs("a", "href")
		
		// now p.introduction a
		// link := e.Attr("href")
		// fmt.Println(link)

		// now p.introduction a strong
		fmt.Println(e.Text)
	})
	c.Visit("https://pythonprogramming.net/parsememcparseface/")
}