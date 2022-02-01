package main

/*

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
		colly.Async(true),
	)


	// how to limit the speed
	c.Limit(&colly.LimitRule{
		RandomDelay: 3 * time.Second,
		// Parallelism: 5,
	})


	// randomize User-Agent
	extensions.RandomUserAgent(c)

	baseURL := "https://books.toscrape.com/catalogue/"

	c.OnHTML("ol.row li", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildAttr("img", "alt"),
			e.ChildText("p.price_color"),
			e.ChildText("p.instock"),
			(baseURL + e.ChildAttr("a", "href")),
		})
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("UserAgent\n", r.Headers.Get("User-Agent"))
		fmt.Println("Visiting", r.URL.String())
	})

	// Set error handler
	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})

	for i := 1; i <= 50; i++ {
		url := fmt.Sprintf("https://books.toscrape.com/catalogue/page-%d.html", i)
		c.Visit(url)
	}

	c.Wait()

	log.Printf("Scraping finished, check file %q for results\n", fName)
}

*/