# colly.Collector.OnHTML
specify a CSS selector (via GoQuery https://github.com/PuerkitoBio/goquery) and a callback function when that tag is reached.

# searching
search for all instances of a tag:
c.OnHTML("a[href]")

search for class
c.OnHTML("div.my-class")

search for id
c.OnHTML("div#my-id")


c.OnHTML("a[href]")


c.OnHTML("a[href]")


# extracting
e.Text
e.Attr()

e.ChildText(<can pass in CSS selectors here, e.g. p a.next>)
e.ChildAttr()
e.ForEach()


# OnRequest
a good place to call r.Abort() if you've hit a certain condition.
numVisited := 0
c.OnRequest(func(r *colly.Request) {
    if numVisited > 100 {
        r.Abort()
    }
    numVisited++
})


# OnResponse
can access the entire HTML document
c.OnResponse(func(r *colly.Response) {
    fmt.Println(r.Body)
})



# helpful links
https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/