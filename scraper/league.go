package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Build struct {
	items []string
	percentages []string
}

func LeagueScraper(url string) {
	collector := colly.NewCollector()
	 
	collector.OnError(func(_ *colly.Response, err error) { 
		log.Println("Something went wrong: ", err) 
	}) 
	 
	// Grab the title stating champion, lane and patch number
	collector.OnHTML("div.info-box", func(element *colly.HTMLElement) { 
		fmt.Printf(element.ChildText("h1") + "\n") 
	})

	// Runes for build
	collector.OnHTML("div.css-e0u3ub", func(element *colly.HTMLElement) {
		fmt.Println("<--- RUNES --->")
		element.ForEach("img", func(i int, h *colly.HTMLElement) {
			if strings.Contains(h.Attr("src"), "grayscale") == false {
				fmt.Println(h.Attr("alt"))
			}
		})
	})

	// Skills priority order
	collector.OnHTML("div.ejiywn42", func(element *colly.HTMLElement) {
		fmt.Println("<--- SKILLS --->")
		skills := make([]string, 0)

		element.ForEach("span", func(i int, h *colly.HTMLElement) {
			skills = append(skills, h.ChildText("strong"))
		})

		fmt.Println(strings.Join(skills, " > "))
	})

	// Itemisation with pick + win percentages
	collector.OnHTML("table.e14sgc0n0", func(table *colly.HTMLElement) {
		caption := table.ChildText("caption")
		fmt.Println("<--- " + strings.ToUpper(caption) + " --->")

		// Split the descriptions as per the table row
		descriptions := make([]Build, 0)
		var build Build

		table.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			// Group the items by current row
			items := make([]string, 0)
			tr.ForEach("img", func(_ int, img *colly.HTMLElement) {
				items = append(items, img.Attr("alt"))
			})

			percentages := make([]string, 0)
			tr.ForEach("strong", func(j int, strong *colly.HTMLElement) {
				percentages = append(percentages, strong.Text)
			})

			build.items = items
			build.percentages = percentages

			// Ignore the header which is always the first <tr>
			if i != 0 {
				descriptions = append(descriptions, build)
			}
		})

		// Could separate logic into function -> not too bad here atm
		for _, desc := range descriptions {
			itemsJoiner := " & "

			if strings.Contains(caption, "Builds") {
				itemsJoiner = " > "
			}

			percentageString := " // "

			for index, percentage := range desc.percentages {
				switch index {
				case 0:
					percentageString = percentageString + "P: " + percentage + " "
				case 1:
					percentageString = percentageString + "W: " + percentage
				}
			}

			fmt.Println(strings.Join(desc.items, itemsJoiner) + percentageString)
		}
	})

	collector.Visit(url)
}