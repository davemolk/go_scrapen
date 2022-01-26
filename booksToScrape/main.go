package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Book stores general information about a book
// type Book struct {
// 	Title string
// 	Link string
// 	Price string
// }

func main() {
	/*
	fName := "books.csv"
	file, err = os.Create(fName)
	if err != nil {
		log.Fatalf("Error: %s. Can't create your csv file.\n", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// write CSV header
	writer.Write([]string{"Name", "Price", "URL"})

	*/
	// create collector
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)
	// extract details
	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		fmt.Println(title)
		link := e.ChildAttrs("div.image_container a", "href")
		fmt.Println(link)
		// price := e.ChildText("div.product_price p.price_color")
		// more like descendants than children? cause this works...
		price := e.ChildText("p.price_color")
		fmt.Println(price)
	})

	c.OnHTML("li.next a", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: \n", r.URL.String())
	})

	c.Visit("https://books.toscrape.com/")
}