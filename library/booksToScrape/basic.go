package main

/*

func basic() {
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

	c := colly.NewCollector()

	c.OnHTML("ol.row li", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildAttr("img", "alt"),
			e.ChildText("p.price_color"),
			e.ChildText("p.instock"),
			e.ChildAttr("a", "href"),
		})
	})
	c.Visit("https://books.toscrape.com/")
	log.Printf("Scraping finished, check file %q for results\n", fName)
}

*/