package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func Scrape(url string) {
	collector := colly.NewCollector()
	 
	collector.OnError(func(_ *colly.Response, err error) { 
		log.Println("Something went wrong: ", err) 
	}) 
	 
	collector.OnHTML("div.info-box", func(element *colly.HTMLElement) { 
		// Grab the title stating champion, lane and patch number
		fmt.Printf(element.ChildText("h1") + "\n") 
	})

	collector.OnHTML("div.css-e0u3ub", func(element *colly.HTMLElement) {
		// Filter for the runes that are "selected", grab their names from the alt tag
		fmt.Println("<--- RUNES --->")
		element.ForEach("img", func(i int, h *colly.HTMLElement) {
			if strings.Contains(h.Attr("src"), "grayscale") == false {
				fmt.Println(h.Attr("alt"))
			}
		})
	})

	collector.Visit(url)
}