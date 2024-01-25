package scraper

import (
	"log"
)

func Scrape(url, selector string) {
	switch selector {
	case "league":
		LeagueScraper(url)
	default:
		log.Fatalf("Error: invalid selector '%s' provided to scraper.", selector)
	}
}