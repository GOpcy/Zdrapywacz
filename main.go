package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Offer struct{
	Title string
	Company string
	Location string
	Experience string
	OperatingMode string
	URL string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("justjoin.it", "www.justjoin.it"),

		colly.CacheDir("./justjoinit_cache"),
	)

	detailCollector := c.Clone()

	//offers := make([]Offer, 0, 200)


	c.OnHTML("div[data-test-id=virtuoso-item-list]", func(e *colly.HTMLElement) {
		
		offers := e.ChildAttrs("a[href]", "href")
		for _, offer := range offers {
			finalURL := e.Request.AbsoluteURL(offer)
			detailCollector.Visit(finalURL)
		}
		
		//fmt.Println(offers)
	})

	c.OnRequest(func(r *colly.Request) {
		//log.Println("visiting", r.URL.String())
	})
	
	detailCollector.OnRequest(func(r *colly.Request) {
		//fmt.Println("visiting", r.URL.String())
	})

	detailCollector.OnHTML(`div.css-tnvghs`, func(e *colly.HTMLElement){
		title := e.ChildText("h1")
		if title == "" {
			log.Println("No title found", e.Request.URL)
		}
		offer := Offer{
			Title: title,
			URL: e.Request.URL.String(),
			Company: e.ChildText("h2"),
			Location: e.ChildText(".css-1o4wo1x"),
		}
		
		e.ForEach(".css-if24yw > div", func(_ int, el *colly.HTMLElement){
			target := el.Text
			switch target{
			case "Experience": 
				offer.Experience = el.DOM.Next().Text()
			case "Operating mode": 
				offer.OperatingMode = el.DOM.Next().Text()
			}	
		})
		fmt.Println(offer.URL," | ", offer.Company, " | ", offer.Location, " | ", offer.Title, " | ", offer.Experience, " | ", offer.OperatingMode);
	})





	c.Visit("https://justjoin.it/job-offers/all-locations?experience-level=junior&orderBy=DESC&sortBy=published&from=0")
}