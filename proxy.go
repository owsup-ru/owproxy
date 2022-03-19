package owproxy

import (
	"log"

	colly "github.com/gocolly/colly"
)

type IP struct {
	IP    string `json:"ip"`
	Port  string `json:"port"`
	City  string `json:"city"`
	Speed string `json:"speed"`
}

func GetList() []IP {
	var ips []IP

	url := "https://hidemy.name/ru/proxy-list/"

	c := colly.NewCollector(
		colly.AllowedDomains("hidemy.name"),
	)

	c.OnHTML("tr", func(e *colly.HTMLElement) {

		var ip = IP{}

		e.ForEach("td", func(i int, e *colly.HTMLElement) {
			log.Println(e.Text, i)

			if i == 0 {
				ip.IP = e.Text
			}

			if i == 1 {
				ip.Port = e.Text
			}

			if i == 2 {
				ip.City = e.Text
			}

			if i == 3 {
				ip.Speed = e.Text
			}
		})

		ips = append(ips, ip)
	})

	c.Visit(url)

	return ips
}
