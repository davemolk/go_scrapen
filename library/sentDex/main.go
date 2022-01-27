package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// sentdex sitemap
func sitemap() {
	// initialize a slice
	knownURLs := []string{}

	c := colly.NewCollector(
		colly.AllowedDomains("pythonprogramming.net"),
	)

	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		knownURLs = append(knownURLs, e.Text)
	})

	c.Visit("https://pythonprogramming.net/sitemap.xml")

	fmt.Println("Collected", len(knownURLs), "URLs")
	for _, url := range knownURLs[:10] {
		fmt.Println(url)
	}
}

// scrape sentdex
func main() {

	sitemap()
	
	c := colly.NewCollector(
		colly.AllowedDomains("pythonprogramming.net"),
	)

	// find and print all links
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	fmt.Println(link)
	// 	// optional visit link
	// 	// c.Visit(e.Request.AbsoluteURL(link))
	// })

	
	// print intro
	// c.OnHTML("p.introduction", func(e *colly.HTMLElement) {
	// 	intro := e.Text
	// 	link := e.ChildAttr("a", "href")
	// 	fmt.Println(intro)
	// 	fmt.Println("read more  at: ", link)
	// })
	

	// print intro to table
	// c.OnHTML("p.introduction + p", func(e *colly.HTMLElement) {
	// 	text := e.Text
	// 	fmt.Println(text)
	// })


	// print programming languages
	// c.OnHTML("div.body ul li", func(e *colly.HTMLElement) {
	// 	lang := e.Text
	// 	fmt.Println(lang)
	// })
	

	// print programming languages using ForEach
	// c.OnHTML("div.body ul", func(e *colly.HTMLElement) {
	// 	e.ForEach("li", func(_ int, el *colly.HTMLElement) {
	// 		lang := el.Text
	// 		fmt.Println(lang)
	// 	})
	// })


	// scrape the table
	// c.OnHTML("table tr", func(e *colly.HTMLElement) {
	// 	e.ForEach("td", func(_ int, el *colly.HTMLElement) {
	// 		row := el.Text
	// 		fmt.Println(row)
	// 	})
	// })

	// c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	// 	language := e.Text
	// 	fmt.Println(language)
	// })

	// c.OnHTML("tr td:nth-of-type(2)", func(e *colly.HTMLElement) {
	// 	points := e.Text
	// 	fmt.Println(points)
	// })

	// c.OnHTML("tr td:nth-of-type(3)", func(e *colly.HTMLElement) {
	// 	kittens := e.Text
	// 	fmt.Println(kittens)
	// })

	// c.OnHTML("table", func(e *colly.HTMLElement) {
	// 	e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
	// 		row := el.Text
	// 		fmt.Println(row)
	// 	})
	// })


	// image
	// c.OnHTML("div.card-content img", func(e *colly.HTMLElement) {
	// 	link := e.Attr("src")
	// 	fmt.Println(link)
	// 	batman := e.Attr("alt")
	// 	fmt.Println(batman)
	// })


	// zen
	// c.OnHTML("pre", func(e *colly.HTMLElement) {
	// 	text := e.Text
	// 	fmt.Println(text)
	// })


	// sitemap
	



	// c.OnHTML("div.body", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.ChildText("p code"))
	// })

	c.Visit("https://pythonprogramming.net/parsememcparseface/")

}