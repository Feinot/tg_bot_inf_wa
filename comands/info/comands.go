package info

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

var (
	msg, water, time string
)

func Inf(city string) string {

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://pogoda.365c.ru/russia/" + city + "/prognoz-na-nedelu"},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("div.content").Each(func(i int, s *goquery.Selection) {

				water = s.Find("td.bw-width-column div.temp").Text()
				time = s.Find("th.thleft p").Text()

			})

		},
	}).Start()
	if time != "" {
		msg = "погода в " + city + " : " + water + " время в " + city + " : " + time
	} else {
		msg = "I don't know such a city"
	}

	time, water = "", ""
	return msg

}
